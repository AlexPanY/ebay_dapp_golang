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
		mapping (address => mapping (bytes32 => Bid)) bids;
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

		Product memory item = Product({
			id: productIndex,
			name: _name, 
			category: _category, 
			imageLink: _imageLink, 
			descLink: _descLink, 
			auctionStartTime: _auctionStartTime, 
			auctionEndTime: _auctionEndTime,
			startPrice: _startPrice, 
			highestBidder: address(0), 
			highestBid: 0, 
			secondHighestBid: 0, 
			totalBids: 0, 
			status: ProductStatus.Open, 
			condition: ProductCondition(_productCondition)
		});

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

	//Bid
	struct Bid {
		address bidder;
		uint productId;
		uint value;
		bool reveable;
	}
	
	//bid
	function bid(uint _productId, bytes32 _bid) payable public returns (bool) {
		Product storage product = stores[productIdInStore[_productId]][_productId];
		
		//??????????????????????????????????????????
		require(block.timestamp >= product.auctionStartTime);
		//??????????????????????????????????????????
		require(block.timestamp <= product.auctionEndTime);

		//???????????????????????????????????????
		require(msg.value > product.startPrice);

		//????????????????????????????????????
		require(product.bids[msg.sender][_bid].bidder == 0);
		
		product.bids[msg.sender][_bid] = Bid(msg.sender, _productId, msg.value, false);
		product.totalBids +=1;

		return true;
	}

	//????????????
	function revealBid(uint _productId, string memory _amount, string memory _secret) public {
		Product storage product = stores[productIdInStore[_productId]][_productId];
		
		require(block.timestamp > product.auctionEndTime);

		bytes32 sealeBid = sha3(_amount, _secret);

		Bid memory bidInfo = product.bids[msg.sender][sealeBid];

		require(bidInfo.bidder > 0);
		require(bidInfo.reveable == false);

		uint refund;

		uint amount  = stringToUint(_amount);

		//?????????????????????????????????????????????????????????????????????
		if(bidInfo.value < amount) {
			refund = bidInfo.value;
		} else {
			if (address(product.highestBidder) == 0) {
				product.highestBidder = msg.sender;
				product.highestBid = amount;
				product.secondHighestBid = product.startPrice;
				refund = bidInfo.value - amount;
			} else {
				//??????????????????????????????????????????????????????
				if (amount > product.highestBid) {
					product.secondHighestBid = product.highestBid;
					product.highestBidder.transfer(product.highestBid);
					product.highestBidder = msg.sender;
					product.highestBid = amount;
					refund = bidInfo.value - amount;
				} else if (amount > product.secondHighestBid) {
					product.secondHighestBid = amount;
					refund = bidInfo.value;
				} else {
					refund = bidInfo.value;
				}
			}
		}

		product.bids[msg.sender][sealeBid].reveable = true;
		if (refund > 0) {
			msg.sender.transfer(refund);
		}
	}
	
	//stringToUint
	function stringToUint(string memory s) private pure returns (uint){
        bytes memory b = bytes(s);
        uint result = 0;
        for(uint i = 0; i < b.length; i++){
            if(b[i] >= 48 && b[i] <= 57){
                result = result * 10 + (uint(b[i]) - 48);
            }
        }
        return result;
    }

	//highestBidderInfo
	function highestBidderInfo(uint _productId) view public returns (address, uint, uint) {
		  Product memory product = stores[productIdInStore[_productId]][_productId];
		  return (product.highestBidder, product.highestBid, product.secondHighestBid);
	}

	//totalBids
	function totalBids(uint _productId) view public returns (uint) {
		Product memory product = stores[productIdInStore[_productId]][_productId];
		return product.totalBids;
	}
}