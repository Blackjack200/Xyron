package com.blackjack200.xyron.nukkit;

import cn.nukkit.Server;
import com.google.common.util.concurrent.ListenableFuture;
import lombok.val;
import lombok.var;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.Future;
import java.util.function.Consumer;

public class BufferedDataFlushPool<T> {
    private final List<Future<T>> futures = new ArrayList<>();
    private final Map<Future<T>, Consumer<T>> callbackMap = new HashMap<>();

    public synchronized void add(ListenableFuture<T> future, Consumer<T> callback) {
        futures.add(future);
        callbackMap.put(future, callback);
    }

    public void poll() {
        val completedFutures = new ArrayList<Future<T>>();
        for (val future : futures) {
            if (future.isDone()) {
                try {
                    val resp = future.get();
                    val consumer = callbackMap.get(future);
                    callbackMap.remove(future);
                    consumer.accept(resp);
                    completedFutures.add(future);
                } catch (Throwable e) {
                    Server.getInstance().getLogger().logException(e);
                }
            }
        }
        futures.removeAll(completedFutures);
    }

    public synchronized void shutdown() {
        var counter = 1 << 16;
        while (this.futures.size() > 0 && counter-- > 0) {
            this.poll();
        }
    }
}
