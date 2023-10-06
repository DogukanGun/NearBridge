use near_sdk::testing_env;
use near_sdk::AccountId;
use near_sdk::VMContext;
use crate::NearBridge;

use near_sdk::PublicKey;
use lazy_static::lazy_static;


lazy_static! {
    static ref MOCK_PUBLIC_KEY: PublicKey = "ed25519:6E8sCci9badyRkXb3JoRpBj5p8C6Tw41ELDZoiihKEtp".parse().unwrap();
}


fn get_context(input: Vec<u8>, is_view: bool) -> VMContext {
    VMContext {
        current_account_id: AccountId::new_unchecked("test.near".to_string()),
        signer_account_id: AccountId::new_unchecked("test.near".to_string()),
        predecessor_account_id: AccountId::new_unchecked("test.near".to_string()),
        input,
        block_index: 0,
        block_timestamp: 0,
        account_balance: 0,
        account_locked_balance: 0,
        storage_usage: 0,
        attached_deposit: 0,
        prepaid_gas: 10u64.pow(18).into(),
        random_seed: Default::default(),
        output_data_receivers: vec![],
        epoch_height: 0,
        view_config: None,
        signer_account_pk:  MOCK_PUBLIC_KEY.clone(),
    }
}

#[test]
fn test_send_message() {
    let context = get_context(vec![], false);
    testing_env!(context);
    let contract = NearBridge::new();
    contract.send_message("Hello, NEAR!".to_string());

    assert!(true, "Function didn't panic");
}