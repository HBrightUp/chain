#![allow(clippy::result_large_err)]

use anchor_lang::prelude::*;
mod instructions;
mod state;
use instructions::*;

declare_id!("3exv41SfQ1kZUJeyYaw78cQ49qY3f2t3duUC8M3cKfiP");

#[program]
pub mod close_account_program {
    use super::*;

    pub fn create_user(ctx: Context<CreateUserContext>, name: String) -> Result<()> {
        create_user::create_user(ctx, name)
    }

    pub fn close_user(ctx: Context<CloseUserContext>) -> Result<()> {
        close_user::close_user(ctx)
    }
}
