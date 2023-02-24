# tirith
An implementation of a Time-based One-Time Password generator as specified in [RFC 4226](https://www.ietf.org/rfc/rfc4226.txt) that mimics [Google Authenticator](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2) on the input (base-32 encoded secret) and output (6 digits).

Requires [go-gtk](https://github.com/mattn/go-gtk/#install).
