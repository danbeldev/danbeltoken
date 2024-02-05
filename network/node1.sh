# bin/bash
geth --datadir network/node1 --networkid 12345 --port 30306 --http --http.port 1111 --allow-insecure-unlock \
  --http.api admin,clique,debug,eth,miner,personal,net,txpool --syncmode full \
  --bootnodes enode://861ca8e860a0580c1f3249fde96cf8483e185234b151ef656d2e62e97408baa47e06b23d79fec5dba9e7571e23f9027c94370bb73c867541f118f7b8fc4357df@127.0.0.1:0?discport=30305