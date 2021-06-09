// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "./@openzeppelin/contracts/access/Ownable.sol";

abstract contract MinterAccessControl is Ownable {
    mapping (address => bool) private _minters;

    event MinterAdded(address indexed minter);
    event MinterRemoved(address indexed minter);

    modifier onlyMinter() {
        require(_minters[_msgSender()], "MinterAccessControl: caller is not the minter");
        _;
    }

    function addMinter(address minter) public onlyOwner {
        require(minter != address(0), "MinterAccessControl: can not add address(0)");
        require(!mintable(minter), "MinterAccessControl: minter already in access list");
        emit MinterAdded(minter);
        _minters[minter] = true;
    }

    function removeMinter(address minter) public onlyOwner {
        require(minter != address(0), "MinterAccessControl: can not add address(0)");
        require(mintable(minter), "MinterAccessControl: minter already not in access list");
        emit MinterRemoved(minter);
        _minters[minter] = false;
    }

    function mintable(address minter) public view returns (bool) {
        return _minters[minter];
    }
}
