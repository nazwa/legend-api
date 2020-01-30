# Golang Legend Client

A very basic and (very) incomplete client to interact with [Legend APIs](https://www.legendware.co.uk/), with a heavy focus on the Contacts API (member creation & pricing)

**This project is not maintained or officially supported by Legend.**

This is a **work in progress** and should be **used at your own risk** :sunglasses:

The HTTP client is currently verbose by default, logging all outgoing requests.
It also offers an ability to persist and reuse responses to most GET requests allowing offline development.

## Dependencies

* github.com/hashicorp/go-retryablehttp
* github.com/shopspring/decimal

There is also a temporary dependency on github.com/spf13/viper, but it's only used by the example app and has nothing to do with the actual library itself.

## License
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
