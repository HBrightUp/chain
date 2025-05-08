#![allow(clippy::result_large_err)]

use anchor_lang::prelude::*;

use instructions::*;

pub mod instructions;
pub mod state;

declare_id!("2gba7f3pD4oLADvnn7hAK7L6kEEERhsesGMVvgTVCWsE");

#[program]
pub mod program_deriveed_addresses {
    use super::*;

    pub fn create_page_visits(ctx: Context<CreatePageVisits>) -> Result<()> {
        create::create_page_visits(ctx)
    }

    pub fn increment_page_visits(ctx: Context<IncrementPageVisits>) -> Result<()> {
        increment::increment_page_visits(ctx)
    }
}
