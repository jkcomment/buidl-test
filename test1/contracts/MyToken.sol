pragma solidity ^0.5.0;

import "./ERC20.sol";
import "./ERC20Detailed.sol";

contract MyToken is ERC20, ERC20Detailed {

    string private _name = "KIM TEST TOKEN";
    string private _symbol = "KTT";
    uint8 private _decimals = 18;

    address account = msg.sender;
    uint value = 100000000000000000000;

    constructor() ERC20Detailed( _name, _symbol, _decimals) public {
        _mint(account, value);
    }
}