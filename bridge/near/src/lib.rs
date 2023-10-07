use near_sdk::{near_bindgen, env};
use near_sdk::borsh::{BorshDeserialize, BorshSerialize};


#[near_bindgen]
#[derive(Default,BorshDeserialize, BorshSerialize,serde::Serialize, serde::Deserialize)]
pub struct NearBridge;

#[near_bindgen]
impl NearBridge {

    pub fn new() -> Self {
        Self
    }

    pub fn send_message(&self, message: String) {
        let sender = env::signer_account_id();
        env::log_str(&format!("Message from {}: {}", sender, message));
    }
}


#[cfg(test)]
mod tests;
