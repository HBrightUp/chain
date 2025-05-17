//! Durable transaction nonces.

pub mod state;
pub use state::State;

// 默认 nonce account pubkey  在交易数据帐户列表的第一个位置，所以它的 index 是0;
pub const NONCED_TX_MARKER_IX_INDEX: u8 = 0;
