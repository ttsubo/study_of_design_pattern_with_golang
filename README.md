# study_of_design_pattern_with_golang

デザインパターンとは、
デザインパターンとは、ソフトウェア開発におけるデザインパターン（型紙（かたがみ）または設計パターン、英: design pattern）とは、過去のソフトウェア設計者が発見し編み出した設計ノウハウを蓄積し、名前をつけ、再利用しやすいように特定の規約に従ってカタログ化したものである。

コンピュータのプログラミングで、素人と達人の間では驚くほどの生産性の差があり、その差はかなりの部分が経験の違いからきている。達人は、さまざまな難局を、何度も何度も耐え忍んで乗り切ってきている。そのような達人たちが同じ問題に取り組んだ場合、典型的にはみな同じパターンの解決策に辿り着く。これがデザインパターンである (GoF)。
それぞれのパターンは、プログラマの間で何度も繰り返し考え出されてきた。したがって、それは最善の解決策ではないかもしれないが、その種の問題に対するトレードオフを考慮した、典型的な解決策ではある。さらに、コストがかかるかもしれない問題解決を実際に行う前の先行調査として、大変役に立つ。パターンに名前が付いていることが重要である。なぜなら、名前が付いていることで問題や解決策を記述したり、会話の中で取り上げたりすることができるようになるからである。
[(以上、ウィキペディア（Wikipedia）より引用)](https://ja.wikipedia.org/wiki/デザインパターン_(ソフトウェア))

## オブジェクトの生成に関するデザインパターンの一覧

|パターン名 | 概要                                                            |
|----------------|----------------------------------------------------------------|
|[Abstract Factory](https://qiita.com/ttsubo/items/48be956e545736fcd3ad)|関連する一連のインスタンスを状況に応じて、適切に生成する方法を提供する。    |
|[Builder](https://qiita.com/ttsubo/items/0331d842bb775a81f12a)         |複合化されたインスタンスの生成過程を隠蔽する。                         |
|[Factory Method](https://qiita.com/ttsubo/items/71dd0e0c0eb35196c32e)  |実際に生成されるインスタンスに依存しない、インスタンスの生成方法を提供する。|
|[Prototype](https://qiita.com/ttsubo/items/9527bac3e4c64acb464c)       |同様のインスタンスを生成するために、原型のインスタンスを複製する。         |
|[Singleton](https://qiita.com/ttsubo/items/2498bf77f81b6ed23198)       |あるクラスについて、インスタンスが単一であることを保証する。              |

## 構造に関するパターンの一覧

|パターン名 | 概要                                                            |
|----------------|----------------------------------------------------------------|
|[Adapter](https://qiita.com/ttsubo/items/35e673ceadff94b258d9)         |元々関連性のない2つのクラスを接続するクラスを作る。|
|[Bridge](https://qiita.com/ttsubo/items/50a329904b4522a958a7)          |クラスなどの実装と、呼出し側の間の橋渡しをするクラスを用意し、実装を隠蔽する。|
|[Composite](https://qiita.com/ttsubo/items/fa0fe5b19bfdf720784a)       |再帰的な構造を表現する。|
|[Decorator](https://qiita.com/ttsubo/items/fa38f236e760c8602e60)       |あるインスタンスに対し、動的に付加機能を追加する。Filterとも呼ばれる。|
|[Facade](https://qiita.com/ttsubo/items/11b0492895a2d4534add)          |複数のサブシステムの窓口となる共通のインタフェースを提供する。|
|[Flyweight](https://qiita.com/ttsubo/items/20afe7747a5d1c90f0a9)       |多数のインスタンスを共有し、インスタンスの構築のための負荷を減らす。|
|[Proxy](https://qiita.com/ttsubo/items/cd0df876aa430c60288d)           |共通のインタフェースを持つインスタンスを内包し、利用者からのアクセスを代理する。Wrapperとも呼ばれる。|

## 振る舞いに関するパターン

|パターン名 | 概要                                                                         |
|-----------------------|----------------------------------------------------------------|
|[Chain of Responsibility](https://qiita.com/ttsubo/items/5eef0f78fa1bd4cf1e56)|イベントの送受信を行う複数のオブジェクトを鎖状につなぎ、それらの間をイベントが渡されてゆくようにする。|
|[Command](https://qiita.com/ttsubo/items/0ad132e8b7009cd4ad95)                |複数の異なる操作について、それぞれに対応するオブジェクトを用意し、オブジェクトを切り替えることで、操作の切替えを実現する。|
|Interpreter            |構文解析のために、文法規則を反映するクラス構造を作る。|
|[Iterator](https://qiita.com/ttsubo/items/78dcf2e18fa4830a18b8)               |複数の要素を内包するオブジェクトのすべての要素に対して、順番にアクセスする方法を提供する。反復子。|
|[Mediator](https://qiita.com/ttsubo/items/0b0c774d19673d7f2cd6)               |オブジェクト間の相互作用を仲介するオブジェクトを定義し、オブジェクト間の結合度を低くする。|
|[Memento](https://qiita.com/ttsubo/items/2c2aabcd4252cd01ca09)                |データ構造に対する一連の操作のそれぞれを記録しておき、以前の状態の復帰または操作の再現が行えるようにする。|
|[Observer](https://qiita.com/ttsubo/items/afb4aa84669e0d3b02c1)               |インスタンスの変化を他のインスタンスから監視できるようにする。Listenerとも呼ばれる。|
|[State](https://qiita.com/ttsubo/items/ace69dc646fafb5cab9a)                  |オブジェクトの状態を変化させることで、処理内容を変えられるようにする。|
|[Strategy](https://qiita.com/ttsubo/items/b09334af3fdcb0186850)               |データ構造に対して適用する一連のアルゴリズムをカプセル化し、アルゴリズムの切替えを容易にする。|
|[Template Method](https://qiita.com/ttsubo/items/a36c6fc9acac3513b4af)        |あるアルゴリズムの途中経過で必要な処理を抽象メソッドに委ね、その実装を変えることで処理が変えられるようにする。|
|[Visitor](https://qiita.com/ttsubo/items/6bd4adc178107445753c)                |データ構造を保持するクラスと、それに対して処理を行うクラスを分離する。|
