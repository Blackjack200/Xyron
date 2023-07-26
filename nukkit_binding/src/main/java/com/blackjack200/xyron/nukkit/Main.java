package com.blackjack200.xyron.nukkit;

import com.github.blackjack200.xyron.AnticheatGrpc;
import com.github.blackjack200.xyron.PlayerOuterClass;
import com.github.blackjack200.xyron.PlayerWrappers;
import com.github.blackjack200.xyron.Xchange;
import io.grpc.ManagedChannelBuilder;
import lombok.val;

import java.util.concurrent.ExecutionException;

public class Main {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        val channel = ManagedChannelBuilder.forAddress("localhost", 8884).usePlaintext().build();
        val client = AnticheatGrpc.newFutureStub(channel).withWaitForReady();
        val req = Xchange.AddPlayerRequest.newBuilder()
                .setPlayer(PlayerOuterClass.Player.newBuilder()
                        .setOsValue(PlayerOuterClass.DeviceOS.Android_VALUE)
                        .setName("IPlayfordev")
                );
        req.putData(0L, Xchange.TimestampedReportData.newBuilder()
                .addData(PlayerWrappers.WildcardReportData.newBuilder().setGameModeData(
                        PlayerWrappers.PlayerGameModeData.newBuilder()
                                .setGameModeValue(PlayerOuterClass.GameMode.Survival_VALUE)
                )).build()
        );
        val ppf = client.addPlayer(req.build());
        while(!ppf.isDone()){
            System.out.println("W");
        }
        val pp = ppf.get();
        System.out.println(pp.getInternalId());
        val rp = Xchange.PlayerReport.newBuilder()
                .setPlayer(pp)
                .setLatency(0.01);
        val jd = client.report(rp.build()).get();
        for (val j : jd.getJudgementsList()) {
            System.out.println(j.getJudgement());
        }
        client.removePlayer(pp).get();
        channel.shutdownNow();
    }
}
