// SPDX-License-Identifier: MIT

pragma solidity ^0.8.6;

contract NFTWorshipEvent {
     event TokenOffered(
        address indexed nft,
        uint256 indexed tokenId,
        address indexed votary,
        uint256 releaseTimestamp,
        address redeemer
    );
    event TokenRedeemed(
        address indexed nft,
        uint256 indexed tokenId,
        address indexed redeemer
    );
}
