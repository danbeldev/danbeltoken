# bin/bash
geth --datadir network/node2 --networkid 12345 --port 30307 --http --http.port 2222 --allow-insecure-unlock \
  --http.api admin,clique,debug,eth,miner,personal,net,txpool --syncmode full \
  --bootnodes enode://861ca8e860a0580c1f3249fde96cf8483e185234b151ef656d2e62e97408baa47e06b23d79fec5dba9e7571e23f9027c94370bb73c867541f118f7b8fc4357df@127.0.0.1:0?discport=30305 \
  --mine --miner.etherbase 0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151 --unlock 0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151 \
  --password network/miner.txt --authrpc.port 8550