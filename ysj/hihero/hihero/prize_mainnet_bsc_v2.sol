// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./interfaces/IHiHeros.sol";

contract Prize is Initializable, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    
    using SafeMath for uint256;

    enum RewardState{ PENDING, SUCCESS, FAIL}

    enum RewardTypes {INVITE, NFT}

    //charge address of project 
    address public constant serviceAddr = 0x35d7749D77773D542aCbf7601216B5cAa5e0004c;

    //asset of invest
    uint256 public constant minValue = 10e18;
    uint256 public constant maxValue = 9999e18;

    // award for project
    uint256 public constant serviceChargeOdds = 10000;

    // four precision  
    uint256 public constant decimal = 1000000;

    // award for superior
    uint256 public constant superiorOdds = 20000;

    // award for NFT owners
    uint256 public constant holderOdds = 18000;

    //rate of winner with claasic mode
    uint256 public constant classicOdds = 10000000;

    //rate of winner with Attributes mode
    uint256 public constant attrOdds = 2500000;

    //rate of winner with singel/double mode
    uint256 public constant singleDoubleOdds = 1800000;

    //rate of winner with small/big mode
    uint256 public constant  smallBigOdds = 1800000;

    //rate of winner with nft holder
    uint256 public constant  nftholderOdds = 18000;

    //quantity of all heros;
    uint256 public constant supplyOfHeros = 13;

    //switch of contract.
    bool private _pause;

    //amount of order list
    uint256 public orderlistIndex;

    //index of latest
    uint256 public orderlistLatest;

    //current section + 1
    uint256  public sectionId;

    // total award of issue.
    uint256 public totalAwardIssued;

    // times of win
    uint256 public numberOfWinnerOrder;

    //award of all users
    uint256 public rewardOfInviteTotals;

    //counts of invitation
    uint256 public countsOfInvited;

    //address of dapp
    address public dapp;

    //address of MATIC coin
    address public  tokenUSDT;

    //information of users
    struct stUserInfo {

        //all award received by user.
        uint256 recvAward;

        //unclaimed award with user
        uint256 unclaimAward;

        // superior address
        address superior;
    }

    // information of order list
    struct stOrderlist{
        
        //current section + 1
        uint256 sectionId;

        //address of bet
        address account;

        //number of bet
        uint256 tokenId;

        // mode of bet
        uint256 mode;

        //asset of bet
        uint256  value;

        // block number on chain
        uint256 blockNum;

        //reward of bet. if win, the amount is base on odds,otherwise is zero.
        uint256  reward;

        /* the state and result of order list 
        0: pending, watting for lottery
        1:winner  
        2： fail
        */
        RewardState    state; 
    }

    //information of lottery
    struct stVictorInfo{
        //winning numbers
        uint256 victorNo;

        //block number
        uint256 blockNum;
    }

    //information of all users.
    mapping(address => stUserInfo) private _users;

    // information of orderlist
    mapping(uint256 => stOrderlist) private _orderlist;

    // victor numbur of history
    mapping(uint256 => stVictorInfo)  private  _victorHistory;

    //Totals number of wins by single heros
    mapping(uint256 => uint256 )  private _winsOfSingleHeros;

    //NFT award by single Heros
    mapping(uint256 => uint256) private _NFTAwardOfSingleHeros;

    //seed for front-dev
    mapping(bytes32 => uint256) private _seedMap;

    //NFT holder
    mapping(uint256 => address) private _holders;

    event Invest(uint256 indexed orderlistIndex, uint8 regionCode);
    event Withdraw(address indexed sender, uint amount);
    event NotifyReward(uint256 indexed sectionId, uint256 indexed victorNo, uint256 blickNum);
    event Deposit(address indexed sender, uint amount, uint balance);
    event PartnerData(address indexed sender, uint256 sectionId, RewardTypes types, address superior, uint256 amount, uint256 timestamp);
    event HeroData(uint256 indexed sectionId, uint256 victorNo, address holder, uint256 awardOfNFT, uint256 totalAward_, uint256 timestamp);

    function initialize() public initializer {

        __Context_init();
        __Ownable_init();
        __ReentrancyGuard_init();
        
        _pause = false;
        sectionId = 1;
        orderlistIndex = 1;

        dapp = 0x30e6c1caC70db1b89a3D7E3b7C201d8617AC28cF;
        tokenUSDT = 0x55d398326f99059fF775485246999027B3197955;
    }

    modifier whenNotPaused() {
        require(!_pause, "pause state");
        _;
    }

    modifier whenPaused() {
        require(_pause, "not pause state");
        _;
    }

    modifier onlyDapp() {
        require(msg.sender == dapp, "only dapp access.");
        _;
    }

    function setDapp(address dapp_) external onlyOwner {
        dapp = dapp_;
    }

    function setUsdt(address usdt_) external onlyOwner {
        tokenUSDT = usdt_;
    }

    function setPause(bool pause_) external onlyOwner {
        _pause = pause_;
    }

    function getPause() external view returns (bool) {
        return _pause;
    }

    function setHolders(uint256[] memory arrIndex_, address[] memory addrHolders_) external whenNotPaused onlyDapp {
        require(arrIndex_.length == addrHolders_.length, "invalid input.");
        require(arrIndex_.length < 13, "too many index input.");

        for(uint256 i = 0; i < arrIndex_.length; i++) {
            require(arrIndex_[i] < 13, "invalid index.");
            _holders[arrIndex_[i]] = addrHolders_[i];
        }
    }

    /**
      * @dev withdraw award for developer
      *
      * Requirements:
      *
      * - `token_` address of ERC20.
      * - `account_` address of receive coin.
      * - `amount_` amount of withdraw
      */ 
    function financialIndependence(address token_, address account_, uint256 amount_) external onlyOwner {
        uint256 asset_ = IERC20Upgradeable(token_).balanceOf(address(this));
        require(amount_ <= asset_, "not enough  asset."); 
        SafeERC20Upgradeable.safeApprove(IERC20Upgradeable(token_), address(this), amount_);
        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(token_), address(this), account_, amount_);
    }

    /**
      * @dev produce a order for input of users.
      *
      * Requirements: 
      *
      * - `tokenId_` hero Id of selected by account.
      * - `mode_` range of actual mode: 0 ~ 20,corresponding with four mode of front-end development.
      * - `asset_` quantity of token bet by account.
      * - `fee_` base fee for platform.
      * - `superior_` the superior address of provided by account. if current account have a superior address, zero address was input.
      * - `fee_` base fee for platform.
           function hash: 0x8a7def87 
      */
    function invest(uint256 tokenId_, uint256 mode_, uint256 asset_, address superior_, bytes32 seed_, uint8 regionCode_) external whenNotPaused nonReentrant returns (bool) {

        require(_seedMap[seed_]  == 0, "duplicate seed.");

        //check input parameter
        require(tokenId_ >= 0 && tokenId_ <= 12, "Invalid tokenId.");
        require(mode_ >= 0 && mode_ <= 20, "Invalid mode.");
        require(asset_ >= minValue && asset_ <= maxValue, " Input vaule out of range.");

        address sender_ = msg.sender;
        require(!AddressUpgradeable.isContract(sender_), "contract not allow.");
        require(sender_ != superior_, "cannot bind self.");

        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), sender_, address(this), asset_.mul(101).div(100));
        _users[serviceAddr].unclaimAward += asset_.div(100);
        
        //award for superior
        _sendSuperiorAmount(superior_, asset_);

        //amount of issue
        _orderlist[orderlistIndex] = stOrderlist(sectionId, sender_, tokenId_, mode_, asset_, block.number, 0, RewardState.PENDING );
        _seedMap[seed_] = orderlistIndex;

        //emit PartnerData(sender_, sectionId, RewardTypes.INVITE, _users[sender_].superior, superiorAmount_, block.timestamp);
        emit Invest(orderlistIndex, regionCode_);

        orderlistIndex++;

        return true;
    }

    /**
      * @dev get lucky number
      *
      * Requirements:
      *
      * - `nonce` a random string given by dapp
      *
      * Returns
      * `uint256` lucky number：0 ~ 12
      */
    function _getRandom(string memory nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encodePacked(msg.sender, nonce, block.timestamp, block.number, block.coinbase))) % supplyOfHeros;
    }
    
    /**
      * @dev  issue award and record information of account.
      *
      * Requirements:
      *
      * - `nonce`  a hash string given by  backend program
      * -  produce a random number and compare with all of orderlist.
      *    this function don't deal with transaction of current block. 
      */
    function notifyReward(string memory nonce) external onlyDapp whenNotPaused nonReentrant  returns (bool) {

        uint256 orderlistIndexTemp_ = orderlistIndex;
        uint256 orderlistLatestTemp_ = orderlistLatest;
        if(orderlistIndexTemp_ == orderlistLatestTemp_ + 1) {
            require(false, "No orderlist exist in trx.");
            return true;
        }

        uint256 odds_ = 0;
        uint256 reward_ = 0;
        bool win = false;  
        uint256 awardOfNFT_;
        uint256 blockNum_ = block.number;

        //record total processed transation.
        uint256 trxCount_;
        uint256 totalAward_;

        uint256 victorNo_ = _getRandom(nonce);
        
        //deal with all orderlist except produed in current block
        for( uint256 start_ = orderlistLatestTemp_ + 1; start_ < orderlistIndexTemp_; start_++) {

            stOrderlist storage list_ = _orderlist[start_];

            //transaction of current block will be ignore.
            if(blockNum_.sub(list_.blockNum) < 1) {
                break;
            }

            if( ++trxCount_ > 50) {
	    	    --trxCount_;
                break;
            }

            awardOfNFT_ += list_.value;

            (win, odds_) = _isWinner(victorNo_, list_.tokenId, list_.mode);

            //fail
            if(!win) {
                list_.state = RewardState.FAIL;
                continue;
            }

            //success
            list_.state = RewardState.SUCCESS;
            reward_ = list_.value.mul(odds_).div(decimal);
            list_.reward = reward_;

            _users[list_.account].unclaimAward += reward_;

            totalAward_ += reward_;
            numberOfWinnerOrder++;
            totalAwardIssued += reward_;
        }

        if(trxCount_ == 0) {
            return true;
        }

        _winsOfSingleHeros[victorNo_] += 1;
        _victorHistory[sectionId] = stVictorInfo(victorNo_, blockNum_);
        emit NotifyReward(sectionId, victorNo_, blockNum_);

        //send award of NFT holder
        awardOfNFT_ = awardOfNFT_.mul(nftholderOdds).div(decimal);
        _sendTokenHolderAmount(victorNo_, awardOfNFT_);
        _NFTAwardOfSingleHeros[victorNo_] += awardOfNFT_;

        address holder_ = getHolderByIndex(victorNo_);

        emit HeroData(sectionId, victorNo_, holder_, awardOfNFT_, totalAward_, block.timestamp);
        emit PartnerData(address(this), sectionId, RewardTypes.NFT, holder_, awardOfNFT_, block.timestamp);

        sectionId++;
        orderlistLatest += trxCount_;

        return true;
    }

    /**
      * @dev withdraw for all account.
      * Requirements:
      *
      * Returns
      * `bool` 
      */
    function withdraw() external  whenNotPaused nonReentrant returns (bool)  {

        uint256 amount_ = _users[msg.sender].unclaimAward;
        if(amount_ == 0) {
            return true;
        }

        uint256 assetOfContract_ = IERC20Upgradeable(tokenUSDT).balanceOf(address(this));
        uint256 issueAward_ = 0;
        if(assetOfContract_ >= amount_) {
            _users[msg.sender].unclaimAward = 0;
            issueAward_ = amount_;
        }
        else {
            _users[msg.sender].unclaimAward = amount_ - assetOfContract_;
            issueAward_ = assetOfContract_;
        }

        SafeERC20Upgradeable.safeApprove(IERC20Upgradeable(tokenUSDT), address(this), issueAward_);
        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), address(this), msg.sender, issueAward_);

        _users[msg.sender].recvAward = _users[msg.sender].recvAward.add(issueAward_);

        emit Withdraw(msg.sender, issueAward_);

        return true;
    }

    /**
      * @dev detail rule with award
      *
      * Requirements:
      *
      * - `victorNo_` lucky number.
      * - `tokenId_`   hero id given by account.
      * - `mode_`   mode: 0 ~ 20
      *   Returns:
      *  `bool`  win or fail
      *  `uint256` the odds of invest.
      */
    function _isWinner(uint256 victorNo_, uint256 tokenId_, uint256 mode_) internal pure returns (bool, uint256) {
        uint256 odds_ = 0;
        bool win = false;
	
	    if(mode_ > 12 && victorNo_ == 0) {
            return (win, odds_);
        }

        //Classic mode
        if (mode_ >= 0 && mode_ <= 12) {
            odds_ = classicOdds;
            if(victorNo_ == tokenId_) {
                win = true;
            }

            return (win, odds_);
        } 
        
        //Fire mode
        if (mode_ == 13) {
            odds_ = attrOdds;
            if(victorNo_ % 4 == 1 &&
              tokenId_ % 4 == 1) {
                win = true;
            }
            return (win, odds_);
        } 

        //Soil mode
        if (mode_ == 14) {
            odds_ = attrOdds;
            if(victorNo_ % 4 == 2 &&
              tokenId_ % 4 == 2) {
                win = true;
            }
            return (win, odds_);
        } 

        //Wind mode
        if (mode_ == 15) {
            odds_ = attrOdds;
            if(victorNo_ % 4 == 3 &&
              tokenId_ % 4 == 3) {
                win = true;
            }
            return (win, odds_);
        } 

        //Water mode
        if (mode_ == 16) {
            odds_ = attrOdds;
            if(victorNo_ % 4 == 0 && 
                victorNo_ != 0 &&
                tokenId_ % 4 == 0 && 
                tokenId_ != 0) {
                win = true;
            }
            return (win, odds_);
        } 

        // Single/Double mode
        if (mode_ == 17 || mode_ == 18) {
            odds_ = singleDoubleOdds;
            if(victorNo_ % 2 == tokenId_ % 2) {
                win = true;
            }
            return (win, odds_);
        } 

        //Small mode
        if (mode_ == 19) {
            odds_ = smallBigOdds;

            if(victorNo_  >= 1 && victorNo_ <= 6 &&
              tokenId_  >= 1 && tokenId_ <= 6) {
                win = true;
            }
            return (win, odds_);
        } 

        // big mode
        if (mode_ == 20) {
            odds_ = smallBigOdds;

            if(victorNo_  >= 7 && victorNo_ <= 12 &&
              tokenId_  >= 7 && tokenId_ <= 12) {
                win = true;
            }
            return (win, odds_);
        } 
        
        return (false, 0);
    }

    /**
      * @dev issue award of NFT holder.
      * Requirements:
      * - `tokenId_` hero Id.
      * - `holderAmount_` amount of award.
      */
    function _sendTokenHolderAmount(uint256 tokenId_, uint256 holderAmount_) internal {
        address holderAddr_ = getHolderByIndex(tokenId_);
        if(holderAddr_ == address(0)) {
            return ;
        }

        _users[holderAddr_].unclaimAward += holderAmount_;

    }

     /**
      * @dev issue award of superior.
      * Requirements:
      * - `tokenId_` hero Id.
      * - `holderAmount_` amount of award.
      */
    function _sendSuperiorAmount(address superior_, uint256 lockedAmount_) internal {
        uint256 superiorAmount_ = lockedAmount_.mul(superiorOdds).div(decimal);
        address sender_ = msg.sender;

        address curSuperior_ = _users[sender_].superior;

        //check bind state of current account 
        if(curSuperior_ == address(0)) {
            _checkRingBind(superior_);
            _users[sender_].superior = superior_;
            _users[superior_].unclaimAward  += superiorAmount_;
        }
        else {
            _users[curSuperior_].unclaimAward += superiorAmount_;
        }

        rewardOfInviteTotals += superiorAmount_;
        countsOfInvited++;

    }

     /**
      * @dev check  bind by circle
      * Requirements:
      * - `superior_` superior address
      */
    function _checkRingBind(address superior_) internal view {
        address sender_  = msg.sender;
        address tempAddr_ = superior_;

        while(true) {
            tempAddr_ = _users[tempAddr_].superior;
            if(tempAddr_ == address(0)) {
                break;
            }

            require(sender_ != tempAddr_, "Ring bind found");
        }
    }

    /**
      * @dev get superior address by provided address
      * Requirements:
      * - `account_` user address
      *
      *  Returns:
      *  `address`  superior address
      */
    function getSuperior(address account_) public view returns (address) {
        return _users[account_].superior;
    }

    /**
      * @dev get some global information 
      * Requirements:
      * 
      *  Returns:
      *  `uint256`  total amount of award: NFT/Invite/IssuedAward
      *  `uint256`  amount of USDT coin in contract.
      */
    function getGlobalInfo() external view returns (uint256, uint256) {

         uint256 awardOfHolder_ = getNFTAward();   
        return (totalAwardIssued.add(awardOfHolder_).add(rewardOfInviteTotals), IERC20Upgradeable(tokenUSDT).balanceOf(address(this)));
    }

    /**
      * @dev get some global information 
      * Requirements:
      * `account_` user address 
      * 
      *  Returns:
      *  `uint256`  user asset 
      */
    function getBalance(address account_) public view returns (uint256) {
        return IERC20Upgradeable(tokenUSDT).balanceOf(account_);
    }

    /**
      * @dev get some global information 
      * Requirements:
      * `orderlistIndex_` index of orderlist
      *
      *  Returns:
      *  `stOrderlist`  information of given orderlist
      */
    function getOrderlistByIndex(uint256 orderlistIndex_ ) external view returns ( stOrderlist memory) {
        return _orderlist[orderlistIndex_];
    }

     /**
      * @dev get some global information 
      * Requirements:
      * `prev_` previous index of orderlist
      * `rear_` rear index of orderlist
      *
      *  Returns:
      *  `stOrderlist[]`  information of given orderlist
      */
    function getOrderlist(uint256 prev_, uint256 rear_) external view returns (stOrderlist[] memory) {
        stOrderlist[] memory list_ = new stOrderlist[](rear_ - prev_);
        uint256 index_ = 0;
        for(uint256 i = prev_; i < rear_; i++) {
            list_[index_++] = _orderlist[i];
        }

        return list_;
    }

     /**
      * @dev get position information of unprocessed orderlist
      * Requirements:
      *
      *  Returns:
      *  `uint256`  index of processed orderlist
      *  `uint256`  index of unprocessed orderlist
      */
    function getorderIndex() external view returns(uint256, uint256) {
        return (orderlistLatest, orderlistIndex);
    }

    /**
      * @dev get information  of hero page
      * Requirements:
      *
      *  Returns:
      *  `uint256`  total award of NFT
      *  `uint256`  index of unprocessed orderlist
      *  `uint256`  index of processed orderlist
      *  `uint256`  index of unprocessed orderlist
      */
    function getHeroPageData() external view returns (uint256, uint256, uint256, uint256) {
        uint256 awardOfHolder_ = getNFTAward();

        return (awardOfHolder_, sectionId - 1, totalAwardIssued,  numberOfWinnerOrder);
    }

    /**
      * @dev get information  of parnter page
      * Requirements:
      *
      *  Returns:
      *  `uint256`  total award of invite
      *  `uint256`  times of issued by invite
      *  `uint256`  total award of invite and NFT holders
      *  `uint256`  times of issued by invite and NFT holders
      */
    function getParnterPageData() external view returns (uint256, uint256, uint256, uint256) {
        uint256 awardOfHolder_ = getNFTAward();   
        return (rewardOfInviteTotals, countsOfInvited, awardOfHolder_.add(rewardOfInviteTotals),  countsOfInvited.add(sectionId - 1));
    }

    /**
      * @dev get total award  of NFT holders
      * Requirements:
      *
      *  Returns:
      *  `uint256`  total award of NFT
      */
    function getNFTAward() public view returns (uint256) {
        uint256 awardOfHolder_ = 0;

        for(uint256 i = 0; i < supplyOfHeros; i++) {
            awardOfHolder_ += _NFTAwardOfSingleHeros[i];
        }

        return awardOfHolder_;
    }

    
    /**
      * @dev get single user address of NFT holders 
      * Requirements:
      *`index_` NFT number
      *
      *  Returns:
      *  `uint256`  user address
      */
    function getHolderByIndex(uint256 index_) public view returns (address) {
        return _holders[index_];
    }

    /**
      * @dev get all user address of NFT holders
      * Requirements:
      *
      *  Returns:
      *  `uint256`  all user address of NFT holders
      */
    function getHolderAll() public view returns (address[] memory) {

        address[] memory holders_ = new address[](supplyOfHeros);

        for(uint i = 0; i < supplyOfHeros; i++) {
            holders_[i] = getHolderByIndex(i);
        }

        return holders_;
    }

    /**
      * @dev get user information by given address
      * Requirements:
      *
      * `account_`  user address
      *  Returns:
      *  `stUserInfo`  user information 
      */
    function getUserInfo(address account_) external view returns (stUserInfo memory) {
        return _users[account_];
    }

    /**
      * @dev get orderlist information by given seed
      * Requirements:
      *
      * `seed_`  the seed of orderlist
      *  Returns:
      *  `stOrderlist`  information of orderlist 
      */
    function getOrderBySeed(bytes32 seed_) external view returns (stOrderlist memory) {
        return _orderlist[_seedMap[seed_]];
    }

    /**
      * @dev get history of lucky number by section id
      * Requirements:
      *
      * `sectionId_`  section Id
      *  Returns:
      *  `stVictorInfo`  information of victorHistory 
      */
    function getVictorHistory(uint256 sectionId_) external view returns (stVictorInfo memory) {
        return _victorHistory[sectionId_];
    }

    /**
      * @dev get the times won by single heros
      * Requirements:
      *
      * `sectionId_`  hero Id
      * Returns:
      *  `uint256`  times won by single heros. 
      */

    function getWinsOfSingleHeros(uint256 heroId_) external view returns (uint256) {
        return _winsOfSingleHeros[heroId_];
    }

    /**
      * @dev get the award won by single heros
      * Requirements:
      *
      * `sectionId_`  hero Id
      *  Returns:
      *  `uint256`  award won by single heros. 
      */
    function getNFTAwardOfSingleHeros(uint256 heroId_) external view returns (uint256) {
        return _NFTAwardOfSingleHeros[heroId_];
    }

    /**
      * @dev get order index by given seed.
      * Requirements:
      *
      * `seed_`  producted by front dev
      *  Returns:
      *  `uint256`  index of order list. 
      */
    function getSeedMap(bytes32 seed_) external view returns (uint256) {
        return _seedMap[seed_];
    }

}
