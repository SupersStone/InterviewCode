// SPDX-License-Identifier: MIT

pragma solidity ^0.8.6;

contract NFTOrderEvent {
    event NewConduit(address user, address conduit);
    event OrderCancelled(address indexed maker, bytes32 indexed orderHash);

    event AllOrdersCancelled(address indexed offerer, uint256 increasedNonce);

    event FixedPriceOrderMatched(
        address indexed maker,
        address indexed taker,
        bytes32 indexed orderHash,
        bytes orderBytes,
        bytes assetsBytes
    );

    event OrderBytesInfo(
        bytes32 order_struct_hash,
        address maker,
        address taker,
        address royalty_recipient,
        uint256 royalty_rate,
        uint64 start_at,
        uint64 expire_at,
        uint64 maker_nonce,
        bool taker_get_nft,
        bytes32 assets_hash
    );

    event AssetBytesInfo(
        bytes32 assets_struct_hash,
        address nft,
        uint256 nft_id,
        uint256 nft_amount,
        address ft,
        uint256 ft_amount
    );
}
