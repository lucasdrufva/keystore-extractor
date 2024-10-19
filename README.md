# Keystore Extractor

Elastic agents use keystores to store credentials and other sensitive information that you don’t want to include directly in your configuration files. Earlier versions of Elastic’s documentation mentioned that "you can use the [beatname] keystore to securely store secrets," which led users to believe that Elastic keystores provide strong security guarantees. 

However, this is not the case—these keystore values are not securely stored, but merely encrypted using a hardcoded default secret. To demonstrate how easily these secrets can be extracted, I have developed this tool, `keystore-extractor`.

After reaching out to the Elastic team, they have updated their documentation in [this pull request](https://github.com/elastic/beats/pull/38667) to clarify that the values in the keystore are obfuscated, not securely encrypted, which more accurately reflects the situation.

## Usage

To build and run the tool:

1. **Build:**
   ```
   go build
   ```

2. **Run:**
   ```
   ./keystore-extractor path-to-keystore
   ```

Keystore files are typically found in locations like `/var/lib/packetbeat/packetbeat.keystore`.

## Security Implications

Elastic currently provides no solution for securely storing secrets in keystores. Users should follow the **principle of least privilege**, ensuring that keystores do not store credentials with excessive privileges. 

Additionally, the use of obfuscation may lead to a false sense of security for users, making it easier to unintentionally expose sensitive data. Instead of improving security, this obfuscation may result in users underestimating the risk of granting access to keystore files, contributing to **security by obscurity**, which is generally ineffective.