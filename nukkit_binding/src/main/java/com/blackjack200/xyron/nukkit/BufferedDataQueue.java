package com.blackjack200.xyron.nukkit;

import com.github.blackjack200.xyron.AnticheatGrpc;
import com.github.blackjack200.xyron.PlayerWrappers;
import com.github.blackjack200.xyron.Xchange;
import com.google.common.util.concurrent.ListenableFuture;
import lombok.val;

import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class BufferedDataQueue {
    private final Map<Long, List<PlayerWrappers.WildcardReportData>> map = new LinkedHashMap<>(80);

    public synchronized void add(long tick, PlayerWrappers.WildcardReportData wdata) {
        if (!map.containsKey(tick)) {
            map.put(tick, new LinkedList<>());
        }
        map.get(tick).add(wdata);
    }

    public synchronized ListenableFuture<Xchange.ReportResponse> flush(AnticheatGrpc.AnticheatFutureStub c, Xchange.PlayerReceipt p, long tick, double latency) {
        val needSend = new LinkedList<Long>();
        for (val k : map.keySet()) {
            if (k <= tick) {
                needSend.add(k);
            }
        }
        needSend.sort(Long::compare);
        val needSendMap = new LinkedHashMap<Long, Xchange.TimestampedReportData>();
        for (val timestamp : needSend) {
            needSendMap.put(timestamp, Xchange.TimestampedReportData.newBuilder()
                    .addAllData(map.get(timestamp)).build());
            map.remove(timestamp);
        }
        val rp = Xchange.PlayerReport.newBuilder()
                .setPlayer(p)
                .setLatency(latency)
                .putAllData(needSendMap)
                .build();
        return c.report(rp);
    }
}