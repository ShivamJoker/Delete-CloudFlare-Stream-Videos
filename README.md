# Cloudflare Stream Video Deletion Program

This program fetchs and deletes all videos from a Cloudflare Stream account using the Cloudflare API.

## Features

- Fetches all videos from a Cloudflare Stream account
- Deletes videos concurrently (up to 10 at a time)
- Cross-platform support (Linux, Windows, macOS)
- Optimized binary size

## Prerequisites

- A Cloudflare account with Stream enabled
- Cloudflare API token with permission to manage Stream

## Installation

Just download the binary from the [release page](https://github.com/ShivamJoker/Delete-CloudFlare-Stream-Videos/releases) for your OS.

> Optionally rename it to something like `delete-cf-videos`

## Usage

Run the program with your Cloudflare API token and account ID:

```sh
./delete-cf-videos --token YOUR_API_TOKEN --accountid YOUR_ACCOUNT_ID
```
Replace `YOUR_API_TOKEN` and `YOUR_ACCOUNT_ID` with your actual Cloudflare API token and account ID.

> [!WARNING]
> All your videos will be gone after you run the program.\
> You won't be able to recover them since CloudFlare doesn't offer any recovery options.


### Getting Your Cloudflare API Token and Account ID

1. **API Token**: 
   - Log in to the [Cloudflare dashboard](https://dash.cloudflare.com/)
   - Go to "My Profile" > "API Tokens"
   - Create a new token with the "Stream:Edit" permission
   - For more information, see [Cloudflare's API token guide](https://developers.cloudflare.com/api/tokens/create)

2. **Account ID**:
   - Log in to the [Cloudflare dashboard](https://dash.cloudflare.com/)
   - Your Account ID is visible in the URL when you're logged in, or in the overview page of your account
   - For more details, see [Cloudflare's documentation on finding your Account ID](https://developers.cloudflare.com/fundamentals/get-started/basic-tasks/find-account-and-zone-ids/)

## Warning

This program will delete ALL videos in your Cloudflare Stream account. Use with caution and ensure you have backups if needed.

## Links

- [Cloudflare Stream Documentation](https://developers.cloudflare.com/stream/)
- [Cloudflare API Documentation](https://developers.cloudflare.com/api/)
- [Go Programming Language](https://golang.org/)

## Support

If you encounter any problems or have any questions, please open an issue in this repository.
