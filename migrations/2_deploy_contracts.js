const EcommerceProduct = artifacts.require("EcommerceProduct");

module.exports = function(deployer) {
  deployer.deploy(EcommerceProduct);
};
