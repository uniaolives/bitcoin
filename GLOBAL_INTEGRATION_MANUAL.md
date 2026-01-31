# Manual de Integração Global - Ecossistema Nomic

Este documento serve como a **Single Source of Truth (SSoT)** para a arquitetura e integração entre os repositórios do ecossistema Nomic.

---

### 1. Camada de Ecossistema
O projeto **Nomic** é uma ponte descentralizada que permite a custódia e transferência segura de Bitcoin para o ecossistema Cosmos. Através de um modelo de custódia descentralizada baseado em multi-assinaturas, o Nomic emite o token `nBTC` (peg 1:1 com BTC), permitindo que o Bitcoin seja utilizado em aplicações DeFi em redes como Osmosis.

---

### 2. Camada de Repositórios
O ecossistema é composto pelos seguintes componentes principais:

*   **Repo A: Bitcoin Core (C++)** (Este repositório)
    *   **Função:** Fork do Bitcoin Core que gerencia a infraestrutura de rede Bitcoin e a custódia das chaves multi-sig dos validadores.
*   **Repo B: Nomic Chain (Go/Cosmos SDK)**
    *   **Função:** Blockchain Proof-of-Stake que gerencia o minting/burning de `nBTC` e a compatibilidade IBC.
*   **Repo C: Relayers (Off-chain)**
    *   **Função:** Processos que monitoram a rede Bitcoin para depósitos e a Nomic Chain para saques, realizando a comunicação entre as redes.
*   **Componente D: Game Server (Distributed Mesh)**
    *   **Função:** Servidor de jogo descentralizado administrado por AGI, utilizando Mesh e Starlink para baixa latência e Bitcoin para ancoragem de estado e economia.

---

### 3. Camada de Fluxo (Bridge Lifecycle)
1.  **Depósito:** O usuário envia BTC para um endereço gerado pela ponte.
2.  **Confirmação:** Relayers detectam o depósito e informam aos validadores da Nomic Chain.
3.  **Minting:** A Nomic Chain emite `nBTC` na conta do usuário.
4.  **IBC Transfer:** O `nBTC` pode ser movido para outras chains Cosmos (ex: Osmosis) via IBC.
5.  **Withdrawal:** O processo inverso queima `nBTC` e libera BTC na rede principal.

---

### 4. Single Source of Truth (SSoT)

| Componente | Repositório Responsável | Consumidores |
| --- | --- | --- |
| **Modelos de Dados (nBTC)** | `nomic-chain-repo` | Relayers, Bitcoin App |
| **Contratos de Interface** | `shared-logic / sample_module.go` | Todos |
| **Configuração de Rede** | `relayer-repo` | Bitcoin & Nomic Nodes |
| **Consenso de Jogo** | `game-server / nomic_game_server.md` | Shard Nodes, AGI Referee |

---

### 5. Regras de Ouro e Guardrails

*   **Consistência de Contrato (API Contract Consistency):** Qualquer mudança no formato de transação multi-sig no Repo A deve ser validada contra os parsers no Repo B e Repo C.
*   **Validação de Versões:** Versões de bibliotecas compartilhadas (como middlewares IBC) devem ser mantidas em sincronia.
*   **Mapeamento de Pontos de Falha:** Identificar onde a latência da rede Bitcoin pode afetar o tempo de minting no Cosmos.
*   **Verificação de Variáveis:** Garanta que as variáveis de ambiente (como `BITCOIN_RPC_URL` e `NOMIC_API`) estejam espelhadas e corretas em todos os serviços.

---

### 6. Fluxo de Trabalho Integrado

1.  **Exploração de Impacto:** Mapear arquivos afetados em todos os repositórios antes de iniciar a codificação.
2.  **Desenvolvimento em Cascata:** Mudanças começam na lógica compartilhada -> Backend (Bitcoin/Nomic) -> Frontend/Relayers.
3.  **Sincronização de Documentação:** Atualizar este manual sempre que uma ponte entre repositórios for alterada.
