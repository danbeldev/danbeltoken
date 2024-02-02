# bin/bash
geth --datadir network/node1 --networkid 12345 --port 30306 --http --http.port 1111 --allow-insecure-unlock \
  --http.api admin,clique,debug,eth,miner,personal,net,txpool --syncmode full \
  --bootnodes enode://a73fde7aab28a8c98ab6741415d3469f65970aaf5f704a85bb51915aed2cca760222f4f25695a275d9383091a1dc39b3627360932e0887b42f5ae91f82b5b194@127.0.0.1:0?discport=30305