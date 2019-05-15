# 動作手順

### 必要なもの
- Geth(Ganacheとganache-cliでは正しく動作しないためローカルでGethを動かす必要あり)
- Goインストール済(Ver 1.11以上)


1. ローカルでGethを起動

2. REMIXから下記のERC20コントラクトをデプロイ
https://remix.ethereum.org

```MyToken.sol
pragma solidity >=0.5.0 <0.7.0;

import "https://github.com/OpenZeppelin/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "https://github.com/OpenZeppelin/openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";

contract MyToken is ERC20, ERC20Detailed {

    string private _name = "KIM TEST TOKEN";
    string private _symbol = "KTT";
    uint8 private _decimals = 18;

    address account = msg.sender;
    uint value = 100000000000000000000;

    constructor() ERC20Detailed( _name, _symbol, _decimals) public {
        _mint(account, value);
    }
}
```

3. ターミナルからプログラムを起動
```Terminal
buidl-test/test1 $ go run main.go -address <デプロイ済みのコントラクトアドレス> 
```

実行結果(起動と同時にERC20トークンの情報を取得し、出力)
```Terminal
2019/05/15 13:58:15 ##### Token info
2019/05/15 13:58:15      ***** name:  KIM TEST TOKEN
2019/05/15 13:58:15      ***** symbol: KTT
2019/05/15 13:58:15      ***** totalSupply:  100000000000000000000
```

送金履歴を取得
REMIXからtransferを実行
from: 0x945Cd603A6754cB13C3D61d8fe240990f86f9f8A
to: 0x450a8a99Bf5ad49dB301F6068C619de2400DE6F7
value: 10000
実行結果
```Terminal
2019/05/15 13:58:47 ##### Transfer event info
2019/05/15 13:58:47      ***** transaction:  0xeb1e0b56b8a770e864d8c479c869e20aec49eac42774b371f5170e7f1e6dbbfe
2019/05/15 13:58:47      ***** from:  0x945Cd603A6754cB13C3D61d8fe240990f86f9f8A
2019/05/15 13:58:47      ***** to:  0x450a8a99Bf5ad49dB301F6068C619de2400DE6F7
2019/05/15 13:58:47      ***** value:  10000

```

