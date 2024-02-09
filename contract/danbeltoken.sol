pragma solidity ^0.8.0;

contract DanBelToken {

    struct User {
        string referralCode;
        string fromReferralCode;
        uint balance;
        uint8 discountPercent;
        string role;
        address addr;
    }

    struct NFT {
        uint id;
        address addr;
        string name;
        string desc;
        uint price;
        uint64 createDateInUnixTimestamp;
        uint count;
    }

    struct NFTCollection {
        uint id;
        address addr;
        string name;
        string desc;
        NFT[] nfts;
    }

    struct Auction {
        uint id;
        uint64 startDate;
        uint64 endDate;
        uint minPrice;
        uint maxPrice;
        address addr;
        uint bet;
        bool isFinal;
        NFTCollection collection;
    }

    NFT[] public nfts;
    NFTCollection[] public nftCollections;
    Auction[] public auctions;

    string public tokenName = "DanBelToken";
    string public symbol = "DanBel";
    uint public decimals = 6;
    uint public totalSupply = 1000000;

    mapping(address => User) public users;
    mapping(string => User) public refCodeUser;
    address[] public userAddress;

    uint public startTime;
    address public owner;

    constructor(address _owner, address tom, address max, address jack) {
        startTime = block.timestamp;
        owner = _owner;

        setUser(owner, 100000, "OWNER");
        setUser(tom, 200000, "BASE_USER");
        setUser(max, 300000, "BASE_USER");
        setUser(jack, 400000, "BASE_USER");
    }

    function setReferral(address addr, string memory code) public view {
        User memory user = users[addr];
        require(keccak256(abi.encodePacked(user.fromReferralCode))!=keccak256(abi.encodePacked("")));
        require(keccak256(abi.encodePacked(user.referralCode))==keccak256(abi.encodePacked(code)));

        User memory refUser = refCodeUser[code];
        require(refUser.discountPercent >= 3);

        user.fromReferralCode = code;
        user.balance += 100;
        refUser.discountPercent += 1;
    }

    function setUser(address addr, uint balance, string memory role) public {
        string memory referralCode = getReferralCode(addr);
        User memory user = User(referralCode, "", balance, 0, role, addr);
        users[addr] = user;
        refCodeUser[referralCode] = user;
        userAddress.push(addr);
    }

    function getReferralCode(address addr) public pure returns (string memory) {
        bytes2 addrBytes = bytes2(bytes20(addr));
        bytes memory str = new bytes(4);
        bytes memory alph = "0123456789abcdef";
        for (uint8 i = 0; i < 2; i++)
        {
            str[i*2] = alph[uint8(addrBytes[i]) / alph.length];
            str[(i*2)+1] = alph[uint8(addrBytes[i]) % alph.length];
        }
        return string(abi.encodePacked("PROFI - ", string(str), "2024"));
    }

    function createNft(
        address addr,
        string memory name,
        string memory desc,
        uint price,
        uint64 createDateInUnixTimestamp,
        uint count
    ) public {
        uint id = nfts.length;
        NFT memory nft = NFT(id, addr, name, desc, price, createDateInUnixTimestamp, count);
        nfts.push(nft);
    }

//    function createNFTCollection(
//        address addr,
//        string memory name,
//        string memory desc,
//        NFT[] memory nft
//    ) public {
//        uint id = nftCollections.length;
//        nftCollections.push(NFTCollection(id, addr, name, desc, nft));
//    }

//    function createAuction(
//        uint collectionId,
//        uint64 startDate,
//        uint64 endDate,
//        uint minPrice,
//        uint maxPrice
//    ) public {
//        uint id = auctions.length;
//        Auction memory action = Auction(
//            id,
//            startDate,
//            endDate,
//            minPrice,
//            maxPrice,
//            address(0),
//            0,
//            false,
//            nftCollections[collectionId]
//        );
//
//        auctions.push(action);
//    }

    function transferNFT(
        uint nftId,
        address from,
        address to,
        uint price
    ) public {
        User memory fromUser = users[from];
        User memory toUser = users[to];

        toUser.balance -= price;
        fromUser.balance += price;

        transferNFT(nftId, to);
    }

    function transferNFT(
        uint nftId,
        address addr
    ) public {
        NFT memory nft = nfts[nftId];
        nft.addr = addr;
        nfts[nft.id] = nft;
    }

//    function auctionBet(
//        uint auctionId,
//        uint bet,
//        address addr
//    ) public {
//        Auction memory auction = auctions[auctionId];
//
//        require(auction.isFinal);
//        require(auction.bet>bet);
//        require(auction.minPrice>bet);
//        require(auction.maxPrice<bet);
//
//        User memory user = users[addr];
//        require(user.balance<bet);
//
//        auction.addr = addr;
//        auction.bet = bet;
//        auctions[auctionId] = auction;
//    }

//    function auctionSetFinal(
//        uint auctionId
//    ) public {
//        Auction auction = auctions[auctionId];
//        auction.isFinal = true;
//        auctions[auctionId] = auction;
//    }
}