# p2pdemo

function：
1. Operator send a purchase to user.
2. user send cheque to storage.(need get a purchase first.)
3. Storage call contract to get money.(For test)
4. Deploy a cash contract into chain, and store the contract address into db.
5. Call the apply cheque method of deployed contract.


Usage:
1. Run 2 instances of p2pdemo in different terminal
2. Copy multiaddress of one instance, paste to another instance's command line, followed by command index. 
For example, to run command 1:
/ip4/127.0.0.1/tcp/10183/p2p/QmSUdCDTDFkX5G4M2yA99iBA5LWMHCpr9NWN4qT11wSdGu 1
Caution: Command 1 must be run first to generate and store a purchase into db, so command 2 can read it.
3. Run command 4 to deploy a cash contract, and store the contract address to db.(Need mine a block to complete.)
4. Run command 5 to call contract method, the contract address has been stored in db.(Need mine a block to complete.)
![支票支付流程图](https://user-images.githubusercontent.com/52232908/131478184-8bbf0137-58d0-493d-b287-3c1b936a84ea.png)


