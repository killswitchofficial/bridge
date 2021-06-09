// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "./@openzeppelin/contracts/access/Ownable.sol";
import "./IFee.sol";

contract FeeFixed is IFee, Ownable {
    uint256 private _fee;

    constructor(uint256 fee_) {
        setFee(fee_);
    }

    function fee() public view returns (uint256) {
        return _fee;
    }

    function setFee(uint256 fee_) public onlyOwner {
        _fee = fee_;
    }

    function calculate(uint256) public view override returns (uint256) {
        return _fee;
    }
}
