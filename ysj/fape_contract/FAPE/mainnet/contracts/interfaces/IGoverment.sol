// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


interface IGoverment {

    enum SwapDirect {
        ENU_GAME_TO_CHAIN,
        ENU_CHAIN_TO_GAME
    }

    //购买入场券
    function buyVip(address account_, address superior_) external returns (uint256);
    //function buyVip(address account_) external returns (uint256);

    //绑定上级
    //function bindSuperior(address superior_) external returns (address);

    //积分兑换成炼金币 
    function FAPEGameToChain(address account_, uint256 amount_) external   returns (bool);

    //炼金币兑换成积分-创建订单
    function FAPEChainToGame( address account_, uint256 amount_) external   returns (bool);

    //用户是否为VIP玩家
    function IsVipPlayer(address account_) external  returns (uint256);

    function getSuperior(address account_) external view returns (address);

    //获取用户向下一层的人数
    function getSubordinateCounts(address account_) external view returns (uint256);

    //用户使用游戏经验包
    function employExpPack(address account_, uint256 amount_) external  returns (bool);

    //获取用户可用的经验包数量
    function getExperiencePackCounts(address account_) external view  returns (uint256);
}



