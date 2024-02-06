# bin/bash
geth --datadir network/node2 --networkid 12345 --port 30307 --http --http.port 2222 --allow-insecure-unlock \
  --http.api admin,clique,debug,eth,miner,personal,net,txpool --syncmode full \
  --bootnodes enode://0279c64187162d82f65d80d755edf1e4d6605a3da38f02bcb196417e5b523f3e17d0b9da59658dbdaa39281015edb9f616a65846d0cd0b3a5fc94d1bde95b13b@127.0.0.1:0?discport=30305 \
  --mine --miner.etherbase 0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151 --unlock 0x3EAb2cDBf65f7F6A507bdA5178cB80F25F5DB151 \
  --password network/miner.txt --authrpc.port 8550