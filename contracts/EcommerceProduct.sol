// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

contract EcommerceProduct {
	enum ProductStatus { Open, Sold, Unsold }
	enum ProductCondition { New, Used }

	uint public productIndex;

	mapping (address => mapping(uint => Product)) stores;
	mapping (uint => address) productIdInStore;

	struct Product {
		uint id;
		string name;
		string category;
		string imageLink;
		string descLink;
		uint auctionStartTime;
		uint auctionEndTime;
		uint startPrice;
		address highestBidder;
		uint highestBid;
		uint secondHighestBid;
		uint totalBids;
		ProductStatus status;
		ProductCondition condition;
	}

	constructor() {
		productIndex = 0;
	}

	function add(
		string memory _name, 
		string memory _category, 
		string memory _imageLink, 
		string memory _descLink, 
		uint _auctionStartTime,
		uint _auctionEndTime, 
		uint _startPrice, 
		uint _productCondition
	) public 
	{
		require (_auctionStartTime < _auctionEndTime);
		productIndex += 1;

		Product memory item = Product(
			productIndex, 
			_name, 
			_category, 
			_imageLink, 
			_descLink, 
			_auctionStartTime, 
			_auctionEndTime,
			_startPrice, 
			address(0), 
			0, 
			0, 
			0, 
			ProductStatus.Open, 
			ProductCondition(_productCondition)
		);

		stores[msg.sender][productIndex] = item;
		productIdInStore[productIndex] = msg.sender;
	}

	//get 
	function get(uint _productId) view public returns (
		uint, 
		string memory, 
		string memory,  
		string memory, 
		string memory, 
		uint, 
		uint, 
		uint, 
		ProductStatus, 
		ProductCondition)
	{
		Product memory item = stores[productIdInStore[_productId]][_productId];

		return (
			item.id, 
			item.name, 
			item.category, 
			item.imageLink, 
			item.descLink, 
			item.auctionStartTime,
			item.auctionEndTime, 
			item.startPrice, 
			item.status, 
			item.condition
		);
	}
}