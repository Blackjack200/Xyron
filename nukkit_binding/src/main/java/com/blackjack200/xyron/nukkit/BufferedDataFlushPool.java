package com.blackjack200.xyron.nukkit;

import cn.nukkit.Server;
import com.google.common.util.concurrent.ListenableFuture;
import lombok.SneakyThrows;
import lombok.val;
import lombok.var;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.Future;
import java.util.function.Consumer;

public class BufferedDataFlushPool {
    private final List<Future<Object>> futures = new ArrayList<>();
    private final Map<Future<Object>, Consumer<Object>> callbackMap = new HashMap<>();

    @SuppressWarnings("unchecked")
    public synchronized <T> void add(ListenableFuture<T> future, Consumer<T> callback) {
        futures.add((Future<Object>) future);
        callbackMap.put((Future<Object>) future, (Consumer<Object>) callback);
    }

    public void poll() {
        val completedFutures = new HashMap<Future<Object>, Consumer<Object>>();
        synchronized (this) {
            for (val future : futures) {
                if (future.isDone()) {
                    completedFutures.put(future, callbackMap.get(future));
                    callbackMap.remove(future);
                }
            }
            futures.removeAll(completedFutures.keySet());
        }
        completedFutures.forEach((future, consumer) -> {
            try {
                consumer.accept(future.get());
            } catch (Throwable e) {
                Server.getInstance().getLogger().logException(e);
            }
        });
    }

    @SneakyThrows
    public synchronized void shutdown() {
        var counter = 1 << 30;
        while (this.futures.size() > 0 && counter-- > 0) {
            this.poll();
            Thread.sleep(100);
        }
    }
}
