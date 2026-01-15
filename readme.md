1️⃣ Struct + Método (básico)

Crie uma struct chamada Pessoa com:

Nome string

Idade int

Implemente um método:

func (p Pessoa) EhMaiorDeIdade() bool


📌 Deve retornar true se idade ≥ 18.

2️⃣ Método com ponteiro

Crie uma struct ContaBancaria com:

Saldo float64

Implemente um método com receiver ponteiro:

func (c *ContaBancaria) Depositar(valor float64)


📌 O saldo deve ser alterado corretamente.

3️⃣ Aninhamento de structs

Crie:

type Endereco struct {
    Cidade string
    Estado string
}

type Cliente struct {
    Nome string
    Endereco
}


📌 Acesse Cidade sem usar cliente.Endereco.Cidade.

4️⃣ Interface simples

Crie uma interface:

type Animal interface {
    Falar() string
}


Implemente a interface para:

Cachorro

Gato

📌 Cada um retorna um som diferente.

5️⃣ Polimorfismo real

Crie uma função:

func EmitirSom(a Animal)


📌 Ela deve imprimir o resultado de a.Falar()
Use com tipos diferentes.

6️⃣ Slice de interface

Crie um slice:

var animais []Animal


📌 Adicione pelo menos dois tipos diferentes que implementem Animal
Percorra o slice e chame Falar().

7️⃣ Interface + Struct embutida

Crie:

type Veiculo interface {
    VelocidadeMaxima() int
}


Crie uma struct base:

type Motor struct {
    Potencia int
}


E uma struct Carro que embuta Motor e implemente Veiculo.

8️⃣ Método que recebe interface

Crie uma função:

func TestarVeiculo(v Veiculo)


📌 Ela deve imprimir a velocidade máxima
Passe tipos diferentes que implementem Veiculo.

9️⃣ Type assertion (assertiva de tipo)

Dada uma interface:

var x interface{}


📌 Atribua um int e depois faça:

type assertion segura (valor, ok := ...)

trate o caso de erro

🔟 Interface vazia + switch de tipos

Crie uma função:

func IdentificarTipo(v interface{})


📌 Use type switch para:

int

string

bool

default