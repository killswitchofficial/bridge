// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "./IBridge.sol";
import "./IWrappedToken.sol";
import "./IFee.sol";
import "./BridgeBase.sol";

contract BridgeBurner is BridgeBase {
    IWrappedToken private _token;

    constructor(IWrappedToken token_, string memory name, IFee fee, ILimiter limiter) BridgeBase(name, fee, limiter) {
        _token = token_;
    }

    function token() public view returns (IWrappedToken) {
        return _token;
    }

    function lock(uint256 amount) external payable override {
        _beforeLock(amount);
        _token.burnFrom(_msgSender(), amount);
        emit Locked(_msgSender(), amount);
    }

    function unlock(address account, uint256 amount, bytes32 hash) external override onlyOwner {
        _setUnlockCompleted(hash);
        _token.mint(account, amount);
        emit Unlocked(account, amount);
    }

    function renounceOwnership() public override onlyOwner {
        _pause();
        Ownable.renounceOwnership();
    }
}
