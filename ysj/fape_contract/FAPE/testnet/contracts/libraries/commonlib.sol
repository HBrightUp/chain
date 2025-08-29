// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library Commonlib{

    //无聊猿空投地址
    address constant   apeAirdrop = 0x79AA4E1cCbbA3fAc824139382b3E4547A197B789;
    
    //国库留库地址
    address constant  funds = 0xe19972fc21e79e408A72fB53829DaF2dcE7c1773;
    
    //游戏奖励发放总地址
    address constant  gameRewardsFunds = 0x858DCCFF98873d81222874bBc56E062432217f35;

    //质押挖矿地址(预留)
    address constant  stakeMint = 0x5c27e1B658ea3fF5126f6aaAF6260121C78498B3;

    //总发行量 100 亿
    uint256 constant totalSupply = 10000000000 * 1e18;

    //用于购买入场券，项目方的 10% 分红地址
    address constant officialUsdt1 = 0xFe286AD60d5F390713d0Ea0679883FC15F4B030b;

    //用于购买入场券，用户没有一级上级时分红到此地址
    address constant officialUsdt2 = 0x0AfB4e0cBD5A74ed5110F4B41911B8012De17b0c;

    //排行榜奖励发放地址
    address constant leaderboard = 0xD9048D0dFC32305F53821D56eE11779b89B0FB95;
    
    //绑定上级的根节点用户地址
    address constant rootSuperior = 0xA1bDf8C186B49Db485Ce3B2B0aD6eb9E8b8D1d23;

}