import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import { HelloSolana } from "../target/types/hello_solana";

describe("example", () => {
  // Configure the client to use the local cluster.
  anchor.setProvider(anchor.AnchorProvider.env());

  const program = anchor.workspace.hello_solana as Program<HelloSolana>;

  it("Is hello!", async () => {
    // Add your test here.
    const tx = await program.methods.hello().rpc();
    console.log("Your transaction signature", tx);
  });
});
