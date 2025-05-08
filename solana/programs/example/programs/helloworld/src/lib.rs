#![allow(clippy::result_large_err)]

use anchor_lang::prelude::*;

declare_id!("DDoyXEJ5ExWYbwEYU4LUKhm1NW1hfAWsLDSP4QtvVz79");

#[program]
pub mod hello_solana {
    use super::*;

    pub fn hello(_ctx: Context<Hello>) -> Result<()> {
        msg!("Hello, Solana!");

        msg!(" Our program id is: {}", &id());

        Ok(())
    }
}

#[derive(Accounts)]
pub struct Hello {} 