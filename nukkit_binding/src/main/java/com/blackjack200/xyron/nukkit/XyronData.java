package com.blackjack200.xyron.nukkit;

import com.github.blackjack200.xyron.Xchange;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public class XyronData {
    @Getter
    private Xchange.PlayerReceipt receipt;
    @Getter
    private BufferedDataQueue queue;
}
