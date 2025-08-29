
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./libraries/commonlib.sol";
import "./interfaces/IGoverment.sol";
import "./interfaces/ITicketNFT.sol";

contract Goverment is IGoverment, Initializable, ReentrancyGuardUpgradeable, OwnableUpgradeable {

    using SafeMath for uint256;

    address constant tokenUSDT = 0xc2132D05D31c914a87C6611C10748AEb04B58e8F;

    address public tokenFAPE;

    address public ticketNFT;

    address public dapp;

    //发放炼金币FAPE的奖励地址
    address public gameAwardAddr;

    //购买礼包所需缴纳的USDT数量 
    uint256 public admissionTicket;

    //是否开启兑换
    bool public isOpenExchange;

    //是否打开邀请榜
    bool public isOpenInviteRanking;

    //是否打开金币消耗榜
    bool public isOpenGoldRanking;

    //具有绑定关系的用户数量 
    uint256 public maxUserIndex;

    //购买礼包赠送的FAPE数量
    uint256 public FAPEGiftAmount;

    //购买过礼包的用户数量
    uint256 public currentUserCounts;

    //用户兑换代币的时间限制
    uint256 public swapTimesnap;

    //金币排行榜榜单数
    uint256 public goldRankingTotalItems;

    //金币排行榜发放奖励的时间间隔
    uint256 public goldRankingTimesnap;

    //金币排行榜已发放奖励的index
    uint256 public goldRankingAwardIndex;

    //最后一次金币排行榜奖励的时间
    uint256 public goldRankingLatestAwardTimestamp;

    //金币排行榜当前期数可发放奖励的数量
    uint256 public goldRankingCurentAward;

    //邀请排行榜榜单数量
    uint256 public inviteRankingTotalItems;

    //邀请排行榜最小更新的时间间隔
    uint256 public inviteRankingTimesnap;

    //最后一次邀请排行榜奖励的时间
    uint256 public inviteRankingLatestAwardTimestamp;

    //邀请榜奖励没有领取的总数量 
    uint256 public inviteRankingAwardWightUnclaim;

    //邀请排行榜当前期数入榜最小值 
    uint256 public inviteRankingMinContribute;

    //邀请排行榜当前期数已有贡献值的用户
    uint256 public inviteRankingIndex;

    //邀请排行榜当前期数用户贡献值总数
    uint256 public inviteRankingContributeTotals;

    //邀请排行榜期数
    uint256 public inviteRankingSection;

    //国库留存的USDT数量
    uint256 public exchequerUsdt;

    struct stInviteRanking {
        address addr;
        uint256 inviteContribute;
    }

    struct stUserInfo {

        //用户的上级
        address superior;

        //下级人数(第一层)
        uint256 subordinateCounts1;

        //下级人数(第二层)
        uint256 subordinateCounts2;

        //兑换的时间戳(Game -> Chain)
        uint256 gameToChain;

         //兑换的时间戳(Chain -> Game)
        uint256 chainToGame;

        //未使用的经验包
        uint256 experiencePack;

        //用户邀请返佣得到的USDT奖励总数
        uint256 awardByInvited;

        //邀请排行榜上，用户上榜并已领取的USDT奖励
        uint256 inviteAwardWithclaimed;

        //邀请排行榜上，用户上榜但未领取的USDT奖励(可累积)
        uint256 inviteAwardWithUnclaimed;

        //邀请贡献值
        uint256 inviteContribute;

        //用户最新邀请的时间戳
        uint256 inviteLatestTimestamp;

        //金币排行榜上，用户已获得的总奖励数
        uint256 awardWithGoldRanking;
    }

    //邀请关系用户列表
    mapping(address => stUserInfo) private _users;

    //邀请关系用户列表遍历index
    mapping(uint256 => address) public _userIndexes;

    //邀请排行榜，默认为前20名
    mapping(uint256 => address) public inviteRankingList;

    //邀请排行榜黑名单，在此列表中的地址不参与邀请排名榜列表
    mapping(address => bool) public inviteRankingBlacklist;

    event BuyVip(address indexed account_, address indexed superior_);
    event BindSuperior(address indexed account_, address indexed superior_);
    event FAPEGameToChainEvent(address indexed account_, uint256 amount_);
    event FAPEChainToGameEvent(address indexed account_, uint256 amount_);

    function initialize() public initializer {
        __ReentrancyGuard_init();
        __Context_init();
        __Ownable_init_unchained();
        
        tokenFAPE = 0xaA3575C28C471539e82049bD2db72063B4Fc852f;
        ticketNFT = 0xb6f951De2fd0F86A21Dcee91BB544D18bEa0d494;
        dapp = 0xA025d787A6D8cF3a8996D7FD415aD96c889F0C52;
        gameAwardAddr = Commonlib.gameRewardsFunds;

        maxUserIndex = 0;
        FAPEGiftAmount = 100000e18;
        currentUserCounts = 0;
        goldRankingAwardIndex = 0;
        goldRankingCurentAward = 0;

        swapTimesnap = 3600 * 24;
        admissionTicket = 200e6;

        goldRankingTotalItems = 100;
        goldRankingTimesnap = 3600 * 24;
        goldRankingLatestAwardTimestamp = block.timestamp;

        inviteRankingTotalItems = 20;
        inviteRankingTimesnap = 3600 * 24;
        inviteRankingMinContribute = 0;
        inviteRankingIndex = 0;
        inviteRankingLatestAwardTimestamp = block.timestamp;
        inviteRankingSection = 1;
        inviteRankingAwardWightUnclaim = 0;
        inviteRankingContributeTotals = 0;

        exchequerUsdt = 0;

        isOpenInviteRanking = false;
        isOpenGoldRanking = false;
        isOpenExchange = true;

        inviteRankingBlacklist[Commonlib.rootSuperior] = true;
    }

    modifier onlyDapp() {
        require(msg.sender == dapp, "only dapp access.");
        _;
    }

    function setTokenFAPE(address tokenFAPE_) external onlyOwner returns (bool) {
        tokenFAPE = tokenFAPE_;
        return true;
    } 

    function setTicketNFT(address ticketNFT_) external onlyOwner returns (bool) {
        ticketNFT = ticketNFT_;
        return true;
    }

    function setDapp(address dapp_) external onlyOwner returns (bool) {
        dapp = dapp_;
        return true;
    }

    function setFAPEGiftAmount(uint256 FAPEGiftAmount_) external onlyOwner returns (bool) {
        FAPEGiftAmount = FAPEGiftAmount_;
        return true;
    }

    function setGameAwardAddr(address gameAwardAddr_) external onlyOwner returns (bool) {
        gameAwardAddr = gameAwardAddr_;
        return true;
    }

    function setSwapTimesnap(uint256 swapTimesnap_) external onlyOwner returns (bool) {
        swapTimesnap = swapTimesnap_;
        return true;
    }

    function setAdmissionTicket(uint256 admissionTicket_) external onlyOwner returns (bool) {
        admissionTicket = admissionTicket_;
        return true;
    }

    function setGoldRankingTotalItems(uint256 rankingTotalItems_) external onlyOwner returns (bool) {
        goldRankingTotalItems = rankingTotalItems_;
        return true;
    }

    function setInviteRankingTotalItems(uint256 rankingTotalItems_) external onlyOwner returns (bool) {
        inviteRankingTotalItems = rankingTotalItems_;
        return true;
    }

    function setGoldRankingTimesnap(uint256 rankingTimesnap_) external onlyOwner returns (bool) {
        goldRankingTimesnap = rankingTimesnap_;
        return true;
    }

    function setInviteRankingTimesnap(uint256 rankingTimesnap_) external onlyOwner returns (bool) {
        inviteRankingTimesnap = rankingTimesnap_;
        return true;
    }
        
    function openExchange(bool isOpenExchange_) external onlyOwner returns (bool) {
        if(isOpenExchange_ != isOpenExchange) {
            isOpenExchange = !isOpenExchange;
        }
    
        return true;   
    }

    function openGoldRanking(bool isOpenGoldRanking_) external onlyOwner returns (bool) {
        if(isOpenGoldRanking_ != isOpenGoldRanking) {
            isOpenGoldRanking = !isOpenGoldRanking;
        }
    
        return true;   
    }

    function openInviteRanking(bool isOpenInviteRanking_) external onlyOwner returns (bool) {
        if(isOpenInviteRanking_ != isOpenInviteRanking) {
            isOpenInviteRanking = !isOpenInviteRanking;
        }
    
        return true;   
    }

    function addInviteRankingBlacklist(address account_) external onlyOwner returns (bool) {
        inviteRankingBlacklist[account_] = true;
        return true;
    }

    function removeInviteRankingBlacklist(address account_) external onlyOwner returns (bool) {
        inviteRankingBlacklist[account_] = false;
        return true;
    }

    function resetRankingConfig() external onlyOwner returns (bool) {
        goldRankingAwardIndex = 0;
        goldRankingAwardIndex = 0;
        goldRankingLatestAwardTimestamp = block.timestamp;
        goldRankingCurentAward = 0;

        return true;
    }

    function buyVip(address account_, address superior_) external override onlyDapp nonReentrant returns (uint256) {
        require(account_ != address(0), " no zero address required.");

        address currentSuperior_ = getSuperior(account_); 
        if(currentSuperior_ == address(0)) {
            require(superior_ != address(0), " no zero address with superior required.");
            require(_bindSuperior(account_, superior_), "bind superior failed.");
            currentSuperior_ = superior_;
        }

        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), account_, Commonlib.officialUsdt1, admissionTicket.mul(15).div(100));
        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), account_, Commonlib.funds, admissionTicket.mul(25).div(100));

        //modify: cancel award of direct and modify award of indirect from 40% to 60%
        //date: 2022-08-25  modifier: hml
        //uint256 directAward = admissionTicket.mul(20).div(100);
        //uint256 indirectAward = admissionTicket.mul(40).div(100);
        uint256 indirectAward = admissionTicket.mul(60).div(100);

        //SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), account_, currentSuperior_, directAward);
        //_users[currentSuperior_].awardByInvited = _users[currentSuperior_].awardByInvited.add(directAward);

        _updateInvateRanking(currentSuperior_, 3);

        address indirect_ = getSuperior(currentSuperior_);
        if (indirect_ != address(0)) {
            SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), account_, indirect_, indirectAward);
            _users[indirect_].awardByInvited = _users[indirect_].awardByInvited.add(indirectAward);
 
            _updateInvateRanking(indirect_, 1);
        }
        else {
            SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), account_, Commonlib.officialUsdt2, indirectAward);
        }
       
        

        uint256 totalVip_ = ITicketNFT(ticketNFT).mintToken(account_, "url");

        if(IERC721Upgradeable(ticketNFT).balanceOf(account_) == 1) {
            currentUserCounts = currentUserCounts + 1;
        }

        if (IERC20Upgradeable(tokenFAPE).balanceOf(gameAwardAddr) >= FAPEGiftAmount) {
            SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenFAPE), gameAwardAddr, account_, FAPEGiftAmount);
        }
        
        _users[account_].experiencePack = _users[account_].experiencePack.add(1);

        emit BuyVip(account_, superior_);
    
        return totalVip_;
    }

    function _bindSuperior(address account_, address superior_) internal  returns (bool) {
        require(account_ != Commonlib.rootSuperior, "rootSuperior cannot bind others.");
        require(account_ != address(0), "sender is zero address. ");
        require(superior_ != address(0), "superior is zero address.");
        require(getSuperior(account_) == address(0), "account bound superior.");
        
        address secondSuperior_ = getSuperior(superior_);
        if(secondSuperior_ == address(0) && superior_ != Commonlib.rootSuperior ) {
            require(false, "superior not bound.");
        }

        _setSuperior(account_, superior_);
        
        emit BindSuperior(account_, superior_);
        return true;
    }

    function FAPEGameToChain(address account_, uint256 amount_) external override  onlyDapp nonReentrant returns (bool) {

        //require( block.timestamp.sub(_users[account_].gameToChain) >= swapTimesnap, "swap once time every day only.");

        require(isOpenExchange, "not open exchange now.");
        require(account_ != address(0), "account address is zero.");
        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenFAPE), gameAwardAddr, account_, amount_.mul(1e18));

        _users[account_].gameToChain = block.timestamp;

        emit FAPEGameToChainEvent(account_, amount_.mul(1e18));
        return true;
    }

    function FAPEChainToGame( address account_, uint256 amount_) external override  onlyDapp nonReentrant returns (bool) {

        //require( block.timestamp.sub(_users[account_].chainToGame) >= swapTimesnap, "swap once time every day only.");
        
        require(isOpenExchange, "not open exchange now.");
        require(account_ != address(0), "account address is zero.");

        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenFAPE), account_, gameAwardAddr, amount_.mul(1e18));

        _users[account_].chainToGame = block.timestamp;
  
        emit FAPEChainToGameEvent(account_, amount_);

        return true;
    }


    function issueGoldRankingAward(address[] memory recipients_, uint256[] memory rate_) external  onlyDapp nonReentrant returns (bool) {

        require(isOpenGoldRanking, "Gold ranking not open.");
        require(!isOpenInviteRanking, "Invite ranking  on running.");

        require(recipients_.length == rate_.length, "length mismatch.");

        if(goldRankingAwardIndex == 0){
            require(block.timestamp.sub(goldRankingLatestAwardTimestamp) >= goldRankingTimesnap, "ranking award once time every day.");
            require(goldRankingCurentAward == 0, "current award no zero");
            goldRankingCurentAward = IERC20Upgradeable(tokenUSDT).balanceOf(Commonlib.leaderboard);
        }
        else {
            require(goldRankingCurentAward > 0, "not enough award issue.");
        }

        goldRankingAwardIndex = goldRankingAwardIndex.add(recipients_.length);
        require(goldRankingAwardIndex <= goldRankingTotalItems, "ranking totals overflow.");

        uint256 awardAsset = 0;
        for(uint i = 0; i < recipients_.length; i++){
            require(rate_[i] <= 10000, "rate overflow.");
            awardAsset = rate_[i].mul(goldRankingCurentAward).div(10000);
            SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), Commonlib.leaderboard, recipients_[i], awardAsset);
            _users[recipients_[i]].awardWithGoldRanking = _users[recipients_[i]].awardWithGoldRanking.add(awardAsset);
        }

        if(goldRankingAwardIndex == goldRankingTotalItems) {
            goldRankingAwardIndex = 0;
            goldRankingCurentAward = 0;
            goldRankingLatestAwardTimestamp = block.timestamp;
        }

        return true;
    }

    function issueInviteRankingAward() external onlyDapp nonReentrant returns (bool) {
        //require(block.timestamp.sub(inviteRankingLatestAwardTimestamp) >= inviteRankingTimesnap, "issue award once time of every day only.");

        /*
        require(!isOpenGoldRanking, "Gold ranking on running.");
        require(isOpenInviteRanking, "Invite ranking not open.");

        uint256 usdtReserve_ = IERC20Upgradeable(tokenUSDT).balanceOf(Commonlib.leaderboard);
        require(usdtReserve_ > exchequerUsdt.add(inviteRankingAwardWightUnclaim), "not enought usdt to issue.");

        usdtReserve_ = usdtReserve_.sub(inviteRankingAwardWightUnclaim).sub(exchequerUsdt);

        exchequerUsdt = usdtReserve_.mul(90).div(100).add(exchequerUsdt);
        usdtReserve_ = usdtReserve_.mul(10).div(100);
        

        address addr;
        for(uint256 i = 0; i < inviteRankingIndex; ++i) {
            addr = inviteRankingList[i];
            _users[addr].inviteAwardWithUnclaimed = _users[addr].inviteContribute.mul(usdtReserve_).div(inviteRankingContributeTotals);
            inviteRankingAwardWightUnclaim = inviteRankingAwardWightUnclaim.add(_users[addr].inviteAwardWithUnclaimed);
            _users[addr].inviteContribute = 0;
        }
        */

        inviteRankingLatestAwardTimestamp = block.timestamp;
        inviteRankingMinContribute = 0;
        inviteRankingIndex = 0;
        inviteRankingContributeTotals = 0;
        inviteRankingSection = inviteRankingSection.add(1);

        return true;
    }

    function claimInviteRankingAward() external nonReentrant returns (bool) {
        require(isOpenInviteRanking, "Invite ranking not open.");

        /*
        uint256 inviteAward_ = _users[msg.sender].inviteAwardWithUnclaimed;
        require( inviteAward_ > 0, "no award of invite ranking to claim.");

        uint256 usdtReserve_ = IERC20Upgradeable(tokenUSDT).balanceOf(Commonlib.leaderboard);
        require( usdtReserve_ >= inviteAward_.add(exchequerUsdt), "usdt not enough to award of invite ranking.");

        SafeERC20Upgradeable.safeTransferFrom(IERC20Upgradeable(tokenUSDT), Commonlib.leaderboard, msg.sender, inviteAward_);
        _users[msg.sender].inviteAwardWithUnclaimed = 0;
        inviteRankingAwardWightUnclaim = inviteRankingAwardWightUnclaim.sub(inviteAward_);
        */
        return true;
    } 

    function _setSuperior(address account_, address superior_) internal   {
        _users[account_].superior = superior_;
        _userIndexes[maxUserIndex++] = account_;
        
        _users[superior_].subordinateCounts1 = _users[superior_].subordinateCounts1.add(1);

        superior_ = getSuperior(superior_);
        if( superior_ != address(0)) {
            _users[superior_].subordinateCounts2 = _users[superior_].subordinateCounts2.add(1);
        }
    }

    function _isInInviteRankingList(address account_) internal view returns (bool) {
        for(uint256 i = 0; i < inviteRankingIndex; ++i) {
            if(inviteRankingList[i] == account_) {
                return true;
            }
        }

        return false;
    }

    function _updateInvateRanking(address account_, uint256 contribue_) internal  {

        if(inviteRankingBlacklist[account_]) {
            return ;
        }

        if(_users[account_].inviteLatestTimestamp < inviteRankingLatestAwardTimestamp) {
            _users[account_].inviteLatestTimestamp = block.timestamp;
            _users[account_].inviteContribute = 0;
        }

        _users[account_].inviteContribute = _users[account_].inviteContribute.add(contribue_);
        
        if(_users[account_].inviteContribute <= inviteRankingMinContribute) {
            return ;
        }


        //inviteRankingContributeTotals = inviteRankingContributeTotals.add(contribue_);

        bool isFind_ = false;
        uint256 minContribute_ = _users[inviteRankingList[0]].inviteContribute;
        uint256 index_ = 0;
        uint256 inviteContrubuteTemp_ = 0;

        for(uint256 i = 0; i < inviteRankingIndex; ++i) {
            if(!isFind_ && inviteRankingList[i] == account_) {
                isFind_ = true;
            }

            inviteContrubuteTemp_ = _users[inviteRankingList[i]].inviteContribute;
            if(inviteContrubuteTemp_ < minContribute_) {
                minContribute_ = inviteContrubuteTemp_;
                index_ = i;
            }
            else if(inviteContrubuteTemp_ == minContribute_ ) {
                if(i == 0 ) {
                    index_ = i;
                    continue;
                }

                if(_users[inviteRankingList[0]].inviteLatestTimestamp < _users[inviteRankingList[i]].inviteLatestTimestamp) {
                    index_ = i;
                }
            }
            else {
                //pass 
            }
        }

        if(_users[account_].inviteContribute < minContribute_) {
            minContribute_ = inviteRankingMinContribute;
        }

        if(inviteRankingIndex == inviteRankingTotalItems ) {
            inviteRankingMinContribute = minContribute_;
        }

        if(!isFind_) {
            if(inviteRankingIndex < inviteRankingTotalItems ) {
                inviteRankingList[inviteRankingIndex++] = account_;
                inviteRankingContributeTotals = inviteRankingContributeTotals.add(contribue_);
            }
            else {
                uint256 oldContribue_ = _users[inviteRankingList[index_]].inviteContribute;
                if(oldContribue_ < contribue_) {
                    inviteRankingContributeTotals = inviteRankingContributeTotals.add(contribue_ - oldContribue_);
                }

                inviteRankingList[index_] = account_;
            }
        } else {
            inviteRankingContributeTotals = inviteRankingContributeTotals.add(contribue_);
        }
    }

    function isSwap(address account_, uint8 direction_) external view returns (bool) {
        require(account_ != address(0), "zero address input.");
        require(direction_ >= 1  && direction_ <= 2, "invalid direction input.");

        bool bRet = false;
        if(direction_ == 1) {
            bRet = block.timestamp.sub(_users[account_].gameToChain) >= swapTimesnap;
        }
        else if (direction_ == 2) {
            bRet = block.timestamp.sub(_users[account_].chainToGame) >= swapTimesnap;
        }
        else {
            // no this condition
        }

        return bRet;
    }


    function IsVipPlayer(address account_) external override  view returns (uint256) {
        return IERC721Upgradeable(ticketNFT).balanceOf(account_);
    }

    function getSuperior(address account_) public view override returns (address) {
        return _users[account_].superior;
    }

    function getSubordinateCounts(address account_) external view override returns (uint256) {
        return _users[account_].subordinateCounts1.add(_users[account_].subordinateCounts2);
    }

    function getSubordinate1Counts(address account_) external view  returns (uint256) {
        return _users[account_].subordinateCounts1;
    }

    function getSubordinate2Counts(address account_) external view  returns (uint256) {
        return _users[account_].subordinateCounts2;
    }

    function employExpPack(address account_, uint256 amount_) external override  onlyDapp returns (bool) {
        _users[account_].experiencePack = _users[account_].experiencePack.sub(amount_);

        return true;
    }

    function getExperiencePackCounts(address account_) external view override returns (uint256) {
        return _users[account_].experiencePack;
    }

    function getInvitedAwardByAccount(address account_) external view returns (uint256) {
        return _users[account_].awardByInvited;
    }

    function getInviteRankingList() external view returns (stInviteRanking[] memory) {

        stInviteRanking[] memory  rankList = new stInviteRanking[](inviteRankingIndex);

        if(inviteRankingIndex == 0) {
            return rankList;
        }

        for(uint256 i = 0; i < inviteRankingIndex; ++i) {
            rankList[i] = stInviteRanking(
                inviteRankingList[i],
                _users[inviteRankingList[i]].inviteContribute
            );
        }
        stInviteRanking memory temp;
        for(uint256 i = 0; i < rankList.length - 1; ++i) {
            for(uint256 j = 0; j < rankList.length - 1 -i; ++j) {
                if(rankList[j].inviteContribute < rankList[j + 1].inviteContribute) {
                    temp = rankList[j];
                    rankList[j] = rankList[j + 1];
                    rankList[j + 1] = temp;
                }
            }
        }
        
        return rankList;
    }

    function getUserInfo(address account_) external view returns (stUserInfo memory ) {
        return _users[account_];
    }

    function getInviteRankingByAccount(address account_) external view returns(uint256, uint256, uint256, uint256, uint256) {
        if(account_ == address(0)) {
            return (inviteRankingSection, inviteRankingLatestAwardTimestamp, 0,0,0);
        }
        else {
            return (inviteRankingSection, inviteRankingLatestAwardTimestamp, 
        _users[account_].inviteContribute, _users[account_].inviteLatestTimestamp, _users[account_].inviteAwardWithUnclaimed);
        }
    }

    function getUSDTAddress() external pure returns (address) {
        return tokenUSDT;
    }
    

}