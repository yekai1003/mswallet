pragma solidity^0.4.25;

import "./mstoken3.sol";
import "./SafeMath.sol";

contract tokenadmin {
    address public admin;
    mstoken3 token;
    constructor(string memory symbol) public {
        admin = msg.sender;
        token = new mstoken3(symbol);
    }
    modifier onlyadmin() {
        require(admin == msg.sender);
        _;
    }
    function upgrade(address addr) public onlyadmin {
        token = mstoken3(addr);
    }
   
    function tokenaddr() public view returns (address) {
        return token.addr();
    }
    
    function mint(address to, uint256 value) public onlyadmin {
        token.mint(to, value);
    }
    function burn(uint256 value) public onlyadmin {
        token.burn(value);
    }
    function transfer(address to, uint256 value) public {
        token.transfer(to, value);
    }
}