// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "./@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "./@openzeppelin/contracts/access/Ownable.sol";
import "./IWrappedToken.sol";
import "./MinterAccessControl.sol";

contract WrappedToken is IWrappedToken, IERC20Metadata, ERC20, ERC20Burnable, Ownable, MinterAccessControl {
    uint8 private _decimals;

    constructor(string memory name, string memory symbol, uint8 decimals_) ERC20(name, symbol) {
        _decimals = decimals_;
    }

    function decimals() public view override(IERC20Metadata, ERC20) returns (uint8) {
        return _decimals;
    }

    function mint(address account, uint256 amount) public override onlyMinter {
        _mint(account, amount);
    }

    function burn(uint256 amount) public override(IWrappedToken, ERC20Burnable) onlyMinter {
        ERC20Burnable.burn(amount);
    }

    function burnFrom(address account, uint256 amount) public override(IWrappedToken, ERC20Burnable) onlyMinter {
        ERC20Burnable.burnFrom(account, amount);
    }
}
