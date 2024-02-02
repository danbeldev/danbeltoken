# bin/bash
geth --datadir network/node2 --networkid 12345 --port 30307 --http --http.port 2222 --allow-insecure-unlock \
  --http.api admin,clique,debug,eth,miner,personal,net,txpool --syncmode full \
  --bootnodes enode://a73fde7aab28a8c98ab6741415d3469f65970aaf5f704a85bb51915aed2cca760222f4f25695a275d9383091a1dc39b3627360932e0887b42f5ae91f82b5b194@127.0.0.1:0?discport=30305 \
  --mine --miner.etherbase 0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151 --unlock 0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151 \
  --password miner.txt --authrpc.port 8550