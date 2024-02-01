pragma solidity ^0.8.0;

contract DanBelToken {

    enum UserRole {
        BASE_USER,
        PUBLIC,
        PRIVATE,
        OWNER
    }

    struct User {
        string referralCode;
        string fromReferralCode;
        uint balance;
        uint8 discountPercent;
        UserRole role;
    }

    string public name = "DanBelToken";
    uint public totalSupply = 1000000;
    uint public cost = 1 ether * 75 / 100000;

    mapping(address => User) public users;
    mapping(string => User) public refCodeUser;
    address[] public userAddress;

    address public owner;

    constructor(address _owner, address investor1, address investor2) {
        owner = _owner;

        setUser(owner, 10000000, UserRole.OWNER);
        setUser(investor1, 300000, UserRole.BASE_USER);
        setUser(investor2, 300000, UserRole.BASE_USER);
    }

    function setReferral(address addr, string code) public {
        User user = users[addr];
        require(user.fromReferralCode == "");

        User refUser = refCodeUser[code];
        require(refUser == user);
        require(refUser.discountPercent >= 3);

        user.fromReferralCode = code;
        user.balance += 100;
        refUser.discountPercent += 1;
    }

    function setUser(address addr, uint balance, UserRole role) public {
        string memory referralCode = getReferralCode(addr);
        User user = User(referralCode, "", balance, 0, role);
        users[addr] = user;
        users[referralCode] = user;
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

    function buy(uint amount) public payable {
        transferFrom_(owner,msg.sender,amount);
    }

    function transferFrom_(address from, address to, uint amount) private {
        users[from].balance -= amount;
        users[to].balance += amount;
    }

    function changeCost(uint newValue) public {
        cost = newValue;
    }
}