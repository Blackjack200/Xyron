<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xchange.proto

namespace GPBMetadata;

class Xchange
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\GPBEmpty::initOnce();
        \GPBMetadata\PlayerWrappers::initOnce();
        \GPBMetadata\Player::initOnce();
        \GPBMetadata\AnticheatTypes::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
xchange.protoxchangeplayer_wrappers.protoplayer.protoanticheat_types.proto"#
PlayerReceipt

internalId (	"B
TimestampedReportData)
data (2.xchange.WildcardReportData"2
ServerExecution
type (	
	extraData (	"�
AddPlayerRequest
player (2.xchange.Player1
data (2#.xchange.AddPlayerRequest.DataEntryK
	DataEntry
key (-
value (2.xchange.TimestampedReportData:8"�
PlayerReport&
player (2.xchange.PlayerReceipt-
data (2.xchange.PlayerReport.DataEntryK
	DataEntry
key (-
value (2.xchange.TimestampedReportData:8"�
JudgementData
type (	%
	judgement (2.xchange.Judgement
message (	1
extraExecutions (2.xchange.ServerExecution"<
ReportResponse*

judgements (2.xchange.JudgementData2�
	Anticheat@
	AddPlayer.xchange.AddPlayerRequest.xchange.PlayerReceipt" @
RemovePlayer.xchange.PlayerReceipt.google.protobuf.Empty" :
Report.xchange.PlayerReport.xchange.ReportResponse" B=
com.github.blackjack200.xyronZxyron/��prokits\\xyron�bproto3'
        , true);

        static::$is_initialized = true;
    }
}

