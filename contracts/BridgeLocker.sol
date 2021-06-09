// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./IBridge.sol";
import "./IFee.sol";
import "./ILimiter.sol";
import "./BridgeBase.sol";

contract BridgeLocker is BridgeBase {
    using SafeERC20 for IERC20;

    IERC20 private _token;

    constructor(IERC20 token_, string memory name, IFee fee, ILimiter limiter) BridgeBase(name, fee, limiter) {
        _token = token_;
    }

    function token() public view returns (IERC20) {
        return _token;
    }

    function lock(uint256 amount) external payable override {
        _beforeLock(amount);
        _token.safeTransferFrom(_msgSender(), address(this), amount);
        emit Locked(_msgSender(), amount);
    }

    function unlock(address account, uint256 amount, bytes32 hash) external override onlyOwner {
        _setUnlockCompleted(hash);
        _token.safeTransfer(account, amount);
        emit Unlocked(account, amount);
    }

    function renounceOwnership() public override onlyOwner {
        uint256 balance = _token.balanceOf(address(this));
        if (balance > 0) {
            _token.safeTransfer(owner(), balance);
        }
        _pause();
        Ownable.renounceOwnership();
    }
}
