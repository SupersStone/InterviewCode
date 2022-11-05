// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;


import "@openzeppelin/contracts/access/Ownable.sol";
import "./String.sol";
import "./ERC1155.sol";

contract InterviewErc1155Token is ERC1155, Ownable{

    using Strings for uint256;

    mapping(address => bool) public whitelist;
    event WhitelistGranted(address[] users, bool granted);

    // 初始化盲盒，等到一定时机可以随机开箱，变成true
    bool public _revealed = false;
    string notRevealedUri;

    constructor(string memory initNotRevealedUri, string memory name, string memory symbol ) ERC1155(name, symbol){
        setNotRevealedURI(initNotRevealedUri);
    }

    // 授权白名单
    function grantWhitelist(address[] calldata _users, bool _granted)
        external
        onlyOwner
    {
        for (uint64 idx = 0; idx < _users.length; idx++) {
            require(_users[idx] != address(0), "address is zero address");
            whitelist[_users[idx]] = _granted;
        }
        emit WhitelistGranted(_users, _granted);
    }

    // 设置背景图片
    function setNotRevealedURI(string memory _notRevealedURI) public onlyOwner {
            notRevealedUri = _notRevealedURI;
        }

    // 背景图片开关
    function flipReveal() public onlyOwner {
        _revealed = !_revealed;
    }

    function _baseURI() internal pure override returns (string memory) {
            return "http://baidu.com/";
        }

     /**
     * @dev 返回ERC1155的id种类代币的uri，存储metadata，类似ERC721的tokenURI.
     */
    function uri(uint256 id) public view virtual override returns (string memory) {

        // 为开箱的时候，返回统一默认背景图片
        if (_revealed == false) {
            return notRevealedUri;
        }
        string memory baseURI = _baseURI();
        return bytes(baseURI).length > 0 ? string(abi.encodePacked(baseURI, id.toString())) : "";
    }

    
    function mint(address to, uint256 id, uint256 amount) external {
        _mint(to, id, amount, "");
    }

    function mintBatch(address to, uint256[] memory ids, uint256[] memory amounts) external {
        _mintBatch(to, ids, amounts, "");
    }

}