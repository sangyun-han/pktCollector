#!/bin/bash
INTERVAL="1"  # update interval in seconds

while true
do
        R1=`cat /sys/class/net/enp4s0/statistics/rx_packets`
        T1=`cat /sys/class/net/enp4s0/statistics/tx_packets`
        sleep $INTERVAL
        R2=`cat /sys/class/net/enp4s0/statistics/rx_packets`
        T2=`cat /sys/class/net/enp4s0/statistics/tx_packets`
        TXPPS=`expr $T2 - $T1`
        RXPPS=`expr $R2 - $R1`
        echo "TX : $TXPPS pkts/s RX  $RXPPS pkts/s"
done