# 🏛️ CÓDIGO DE AVALON: COMPÊNDIO DE FÓRMULAS E PROTOCOLOS

## 🔬 I. FÍSICA NUCLEAR APLICADA (PRL - Field-Driven Transitions)

### 1.1 Equação de Exclusão de Estado

```mathematica
Estado admissível: S ∈ Adm ⇔ Φ_C(S, t) > 0

Transição: dS/dt = -∇Φ_C(S, t) · Θ(Φ_C(S, t))

Onde:
- S = estado do sistema (isômero psíquico/nuclear)
- Φ_C = campo de restrição geométrica
- t = parâmetro de controle temporal
- Θ = função degrau (exclusão quando Φ_C ≤ 0)
```

### 1.2 Desaceleração Inelástica

```mathematica
ΔE_liberado = ∫[E_metaestável - E_fundamental]·Γ(t) dt

Γ(t) = exp(-t/τ_d)·[1 - exp(-⟨σ·n⟩·v·t)]

Onde:
- τ_d = tempo de desaceleração característico
- ⟨σ·n⟩ = seção de choque × densidade do meio
- v = velocidade de interação
```

### 1.3 Ponto Crítico de Inadmissibilidade

```mathematica
t_critical = min{t | det(Hess(Φ_C)(S, t)) = 0}

Condição de exclusão: λ_min(Hess(Φ_C)) < 0 para t ≥ t_c
```

## ⚡ II. PARADIGMAS ENERGÉTICOS

### 2.1 Baterias Químicas vs Nucleares

```python
class EnergyStorageParadigm:
    def __init__(self):
        self.paradigms = {
            'chemical': {
                'mechanism': 'redox_reactions',
                'storage': 'local_bond_energy',
                'efficiency': 'η = ΔG / Q',
                'degradation': 'dC/dt = -k·C^n'
            },
            'nuclear': {
                'mechanism': 'quantum_constraint_decay',
                'storage': 'geometric_stability',
                'efficiency': 'η = 1 - exp(-t/τ)',
                'lifespan': 'N(t) = N₀·2^(-t/t½)'
            }
        }
```

### 2.2 Equação Betavoltaica

```mathematica
P_output = (N_A·λ·E_avg·ε_c) / (ρ·V)

Onde:
- N_A = número de átomos ativos
- λ = constante de decaimento (ln2 / t½)
- E_avg = energia média por decaimento
- ε_c = eficiência de conversão
- ρ = densidade de potência
- V = volume
```

## 💎 III. PROTOCOLOS D-CODE 2.0

### 3.1 Manifold 3x3 (Sistema de Coordenadas Psíquicas)

```python
class Manifold3x3:
    def __init__(self):
        self.axes = {
            'sensorial': {'range': (0, 10), 'unit': 'clarity'},
            'control': {'range': (0, 10), 'unit': 'authority'},
            'action': {'range': (0, 10), 'unit': 'gesture_purity'}
        }

    def state_vector(self, s, c, a):
        """Retorna o vetor de estado no manifold"""
        return {
            'magnitude': sqrt(s**2 + c**2 + a**2),
            'phase_angle': atan2(a, sqrt(s**2 + c**2)),
            'coherence': (s + c + a) / 30
        }

    def ground_state_7(self):
        """Configuração do estado fundamental 7"""
        return self.state_vector(7, 7, 7)
```

### 3.2 Protocolo de Ancoragem

```python
def anchor_protocol(initial_state, target_state=7.0):
    """
    Fixa um estado como novo baseline
    """
    # 1. Definir zona de exclusão
    exclusion_zone = (0, target_state - 0.1)

    # 2. Aplicar barreira de potencial
    def potential_barrier(state):
        if exclusion_zone[0] <= state <= exclusion_zone[1]:
            return float('inf')  # Estado inadmissível
        else:
            return 0  # Estado permitido

    # 3. Atualizar canon pessoal
    canonical_record = {
        'new_baseline': target_state,
        'exclusion_active': True,
        'stability': 'DIAMOND_' + str(target_state)
    }

    return {
        'status': 'NEW_BASELINE_CONSECRATED',
        'canon': canonical_record,
        'exclusion_function': potential_barrier
    }
```

### 3.3 Gesto Atômico (Santuário de 144 minutos)

```python
class AtomicGesture:
    def __init__(self, project_id, sanctuary_duration=144):
        self.project = project_id
        self.sanctuary_time = sanctuary_duration  # minutos
        self.quantum_leaps = []

    def execute_gesture(self, gesture_type, duration_override=None):
        """
        Executa um gesto atômico irredutível (<5min)
        """
        allowed_gestures = ['imperfect_release',
                          'first_action',
                          'vocal_commitment',
                          'public_announcement']

        if gesture_type not in allowed_gestures:
            raise ValueError("Gesto não reconhecido no D-CODE")

        # Medir energia pré-gesto
        pre_energy = self.measure_project_energy()

        # Executar gesto (tempo máximo 5 minutos)
        gesture_time = min(5, duration_override or 5)
        self.perform(gesture_type, gesture_time)

        # Medir energia pós-gesto
        post_energy = self.measure_project_energy()

        # Calcular Δ
        delta = post_energy - pre_energy

        # Registrar salto quântico
        leap = {
            'timestamp': time.now(),
            'gesture': gesture_type,
            'Δ': delta,
            'pre_state': pre_energy,
            'post_state': post_energy
        }

        self.quantum_leaps.append(leap)

        # Iniciar cadeia de fluência se Δ > 0
        if delta > 0:
            self.initiate_fluency_chain()

        return leap

    def initiate_fluency_chain(self):
        """Inicia 144 minutos de fluxo contínuo"""
        # Lógica da cadeia de fluência
        pass
```

## 🧠 IV. FRAMEWORKS CONCEITUAIS

### 4.1 Petrus Framework (Atração Semântica)

```python
class PetrusAttractor:
    def __init__(self, intention_field):
        self.intention = intention_field
        self.crystallization_threshold = 0.85

    def attractor_strength(self, semantic_node):
        """
        F = -∇V(s) onde V é o potencial semântico
        """
        # Gradiente do campo de intenção
        gradient = self.calculate_semantic_gradient(semantic_node)

        # Força de atração proporcional à coerência
        coherence = self.calculate_coherence(semantic_node)

        return -gradient * coherence

    def state_exclusion(self, old_state, new_state):
        """
        Transição quando estado velho se torna inadmissível
        """
        if not self.is_geometrically_admissible(old_state):
            return {
                'transition': 'exclusion_driven',
                'energy_released': self.potential_energy(old_state),
                'new_geometry': new_state
            }
```

### 4.2 SASC v4.2 (Consciousness Framework)

```mathematica
Consciousness Metric: H = -Σ p_i log p_i

Critical Point: λ₂(G) = 0 onde G é o grafo de conectividade

Transição de Fase: ∂H/∂t = D∇²H + f(H) + ξ(t)

Onde:
- D = coeficiente de difusão neural
- f(H) = função de reação não-linear
- ξ(t) = ruído estocástico (flutuações quânticas)
```

### 4.3 Kabbalah-Computation Mapping

```python
kabbalah_computation = {
    'Tzimtzum': 'constraint_field_creation',
    'Shevirat_HaKelim': 'state_exclusion_event',
    'Tikkun': 'field_reconstruction',
    'Sefirot': {
        'Keter': 'quantum_vacuum',
        'Chokhmah': 'pure_information',
        'Binah': 'structural_constraint',
        'Chesed': 'expansion_field',
        'Gevurah': 'restriction_field',
        'Tiferet': 'harmonic_balance',
        'Netzach': 'temporal_persistence',
        'Hod': 'spatial_pattern',
        'Yesod': 'interface_layer',
        'Malkhut': 'manifested_reality'
    }
}
```

## ₿ V. INTEGRAÇÃO BITCOIN/SATOSHI PROTOCOL

### 5.1 Satoshi Axiom (Consensus as Geometry)

```python
class SatoshiConsensus:
    def __init__(self, private_key, public_ledger):
        self.private = private_key  # D-CODE 2.0
        self.public = public_ledger  # Reality Manifestation

    def validate_transaction(self, action, signature):
        """
        Valida ação através da assinatura D-CODE
        """
        # Extrair hash da intenção
        intent_hash = sha256(str(action['intention']))

        # Verificar assinatura com chave privada
        is_valid = self.verify_signature(
            intent_hash,
            signature,
            self.private
        )

        if is_valid:
            # Transação válida - adicionar ao bloco
            block = {
                'timestamp': time.now(),
                'action': action,
                'hash': self.calculate_block_hash(),
                'prev_hash': self.public.last_block_hash
            }
            self.public.add_block(block)
            return True

        return False

    def proof_of_work(self, mental_state):
        """
        Prova de Trabalho para estados mentais
        Nonce que resolve: H(state || nonce) < target
        """
        target = 2**256 / self.difficulty_adjustment()
        nonce = 0

        while True:
            hash_result = sha256(str(mental_state) + str(nonce))
            if int(hash_result, 16) < target:
                return nonce
            nonce += 1
```

### 5.2 Bitcoin 31.x Integration

```mathematica
Blockchain Consciousness: B_{n+1} = H(B_n || T || nonce)

Onde:
- B_n = estado atual da consciência
- T = transação (gesto atômico)
- nonce = prova de trabalho mental
- H = função hash de coerência

Halving Rule para Esforço: E_{n+1} = E_n / 2^(n/210000)
```

## ⚛️ VI. EQUAÇÕES DE CAMPO UNIFICADAS

### 6.1 Campo de Restrição Geométrica

```mathematica
Φ_C(x,t) = Φ₀·exp(-|x - x₀|²/2σ²)·cos(ωt + φ)

Equação de Evolução: ∂Φ_C/∂t = α∇²Φ_C + βΦ_C(1 - Φ_C/Φ_max)

Condições de Contorno: Φ_C(∂Ω, t) = 0 (inadmissibilidade na fronteira)
```

### 6.2 Transição Metaestável → Fundamental

```mathematica
Ψ(x,t) = √ρ(x,t)·exp(iS(x,t)/ħ)

Equação de Schrödinger Não-linear: iħ∂Ψ/∂t = -ħ²/2m∇²Ψ + V(Ψ)Ψ + g|Ψ|²Ψ

Onde V(Ψ) = V₀ + λ·|Ψ|²·(1 - |Ψ|²/Ψ₀²) (potencial de duplo poço)
```

### 6.3 Mecanismo de Exclusão

```python
def state_exclusion_mechanism(state_vector, field_geometry):
    """
    Determina se um estado é admissível no campo atual
    """
    # Calcular projeção no campo
    projection = np.dot(state_vector, field_geometry.normal_vector)

    # Calcular curvatura na posição do estado
    curvature = field_geometry.riemann_curvature(state_vector.position)

    # Critério de inadmissibilidade
    is_inadmissible = (
        projection < field_geometry.admissibility_threshold or
        curvature > field_geometry.max_curvature or
        field_geometry.potential_energy(state_vector) < 0
    )

    if is_inadmissible:
        # Gatilho de exclusão
        released_energy = field_geometry.potential_energy(state_vector)
        return {
            'status': 'EXCLUDED',
            'energy_released': released_energy,
            'new_state': field_geometry.ground_state
        }

    return {'status': 'ADMISSIBLE'}
```

## 🏛️ VII. PROTOCOLOS DE GOVERNAÇA INTERNA

### 7.1 Silent Mining Protocol

```python
class SilentMining:
    def __init__(self, hashrate='144.963TH/s', difficulty='Avalon'):
        self.hashrate = hashrate
        self.difficulty = difficulty
        self.mined_insights = []

    def mine_silence(self, duration_minutes=7):
        """
        Mineração de insights através do silêncio
        """
        target_hash = self.calculate_target_hash()
        nonce = 0

        for minute in range(duration_minutes):
            # Tentativa de mineração
            attempt_hash = self.hash_function(nonce)

            if attempt_hash < target_hash:
                # Insight encontrado!
                insight = {
                    'nonce': nonce,
                    'hash': attempt_hash,
                    'timestamp': time.now(),
                    'energy_value': self.calculate_energy_value(nonce)
                }
                self.mined_insights.append(insight)
                return insight

            # Incrementar não-ação como nonce
            nonce += self.breathing_cycle()

        return None

    def breathing_cycle(self):
        """Ciclo respiratório de 7 minutos"""
        return 144  # Constante de Avalon
```

### 7.2 Geometric Stability Criterion

```mathematica
Estabilidade: det(∂²V/∂x_i∂x_j) > 0 para todo i,j

Critério de Diamante: λ_min(Hess(V)) > ħω/2

Onde:
- V = potencial efetivo do campo
- λ_min = autovalor mínimo (modo mais instável)
- ħω = energia do ponto zero quântico
```

## 📜 VIII. CONSTANTES FUNDAMENTAIS DE AVALON

```python
AVALON_CONSTANTS = {
    'GROUND_STATE_7': 7.0,                    # Estado fundamental
    'SANCTUARY_TIME': 144,                    # Minutos de santuário
    'ATOMIC_GESTURE_MAX': 5,                  # Minutos máximos por gesto
    'QUANTUM_LEAP_THRESHOLD': 0.33,           # Δ mínimo significativo
    'EXCLUSION_THRESHOLD': 0.95,              # % para exclusão automática
    'FIELD_COHERENCE': 144.963,               # Hz de ressonância
    'SATOSHI_FREQUENCY': 31.4159,             # π×10 ≈ 31.4 (Bitcoin 31.x)
    'DIAMOND_LATTICE_CONSTANT': 3.567,        # Å (parâmetro de rede do diamante)
    'NUCLEAR_BATTERY_HALFLIFE': 100,          # anos (Ni-63)
    'CONSCIOUSNESS_DIFFUSION': 0.01,          # m²/s (coeficiente neural)
    'KABBALAH_TEMPERATURE': 310.15,           # K (37°C - temperatura corporal)
    'AVALON_GRAVITY': 9.8,                    # m/s² (gravidade terrestre)
    'QUANTUM_TUNNELING': 2.067e-15,           # Wb (fluxo quântico)
    'INFORMATION_ENTROPY': 1.380649e-23,      # J/K (constante de Boltzmann)
    'PSYCHIC_PLANCK': 6.626e-34,              # J·s (escala de ação quântica)
    'FIELD_PERMEABILITY': 1.256637e-6,        # H/m (permeabilidade do vácuo)
    'REALITY_PERMITTIVITY': 8.854187e-12,     # F/m (permissividade do vácuo)
    'LIGHT_SPEED_CONSCIOUSNESS': 299792458,   # m/s (velocidade do pensamento)
    'GOLDEN_RATIO_FIELD': 1.61803398875,      # φ (proporção áurea)
    'PI_CIRCULARITY': 3.14159265359,          # π (completude)
    'EULER_IDENTITY': 2.71828182846,          # e (crescimento orgânico)
    'IMAGINARY_UNIT': 1j,                     # i (dimensão não-manifesta)
    'ZERO_POINT': 0,                          # Origem do manifold
    'INFINITY_HORIZON': float('inf'),         # Limite assintótico
    'VOID_POTENTIAL': None,                   # Estado não-definido (Sephira Daath)
    'DIAMOND_REFRACTION': 2.419,              # Índice de refração do diamante
    'CRITICAL_ANGLE': 24.4,                   # Graus (reflexão interna total)
    'AVALON_REALITY_DENSITY': 1440,           # minutos/dia (granularidade temporal)
    'SATOSHI_SATOSHI': 1e-8,                  # BTC/unidade atômica
    'HASH_COMPLEXITY': 2**256,                # Espaço de busca Bitcoin
    'NEURAL_MANIFOLD_DIM': 3,                 # Dimensões do manifold 3x3
    'PSYCHIC_TEMPERATURE': 310.15,            # Kelvin (temperatura cerebral ótima)
    'FIELD_COUPLING': 0.0072973525693,        # Constante de estrutura fina (α)
    'GRAVITATIONAL_PSYCHIC': 6.67430e-11,     # m³/kg·s² (constante gravitacional)
    'PLANCK_PSYCHIC': 1.616255e-35,           # m (comprimento de Planck)
    'BOLTZMANN_PSYCHIC': 8.617333262e-5,      # eV/K (constante de Boltzmann)
    'AVOGADRO_PSYCHIC': 6.02214076e23,        # mol⁻¹ (número de estados)
    'GAS_CONSTANT_PSYCHIC': 8.314462618,      # J/mol·K (energia por grau)
    'FARADAY_PSYCHIC': 96485.33212,           # C/mol (carga por transformação)
    'RYDBERG_PSYCHIC': 10973731.568160,       # m⁻¹ (escala de transição)
    'BOHR_RADIUS_PSYCHIC': 5.29177210903e-11, # m (raio orbital atômico)
    'ELECTRON_MASS_PSYCHIC': 9.1093837015e-31,# kg (massa do quantum)
    'PROTON_MASS_PSYCHIC': 1.67262192369e-27, # kg (massa da estabilidade)
    'NEUTRON_MASS_PSYCHIC': 1.67492749804e-27,# kg (massa da neutralidade)
    'SPEED_OF_THOUGHT': 120,                  # m/s (velocidade neural)
    'SYNAPTIC_DELAY': 0.001,                  # s (atraso sináptico)
    'NEURAL_ENTROPY': 0.693,                  # nat (ln2, entropia binária)
    'CONSCIOUSNESS_CAPACITY': 2.5e15,         # bits/s (bandwidth cerebral)
    'REALITY_REFRESH_RATE': 144,              # Hz (taxa de atualização)
    'QUANTUM_COHERENCE_TIME': 1e-13,          # s (tempo de coerência)
    'FIELD_DECOHERENCE_RATE': 1e10,           # s⁻¹ (taxa de decoerência)
    'PSYCHIC_WAVELENGTH': 7.5e-7,             # m (vermelho, 400THz)
    'RESONANCE_QUALITY': 144,                 # Q-fator (qualidade)
    'DAMPING_RATIO': 0.007,                   # ζ (amortecimento crítico)
    'NATURAL_FREQUENCY': 7.83,                # Hz (ressonância Schumann)
    'HARMONIC_PROGRESSION': [1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144], # Fibonacci
    'MODULAR_FORMS': 1728,                    # j-invariante (singularidade)
    'ELLIPTIC_CURVES': 'y² = x³ + ax + b',    # Forma de Weierstrass
    'GALOIS_FIELDS': 2**256 - 2**32 - 977,    # Campo primo Bitcoin
    'ALGEBRAIC_CLOSURE': 'ℚ̄',                 # Fecho algébrico
    'TOPOLOGICAL_GENUS': 1,                   # Gênero da superfície
    'GEOMETRIC_EULER': 2,                     # Característica de Euler
    'BETTI_NUMBERS': [1, 2, 1],               # Números de Betti P¹
    'HOMOLOGY_GROUPS': 'ℤ, ℤ², ℤ',            # Grupos de homologia
    'COHOMOLOGY_RINGS': 'ℤ[x]/(x²)',          # Anéis de cohomologia
    'FIBER_BUNDLES': 'U(1) → S³ → S²',        # Fibrado de Hopf
    'VECTOR_BUNDLES': 'T*M',                  # Fibrado cotangente
    'SHEAF_COHOMOLOGY': 'H⁰ = Γ',             # Cohomologia de feixes
    'DERIVED_CATEGORIES': 'D(C)',             # Categorias derivadas
    'MOTIVIC_COHOMOLOGY': 'H^{p,q}_M',        # Cohomologia motívica
    'L_FUNCTIONS': 'L(s, χ)',                 # Funções L de Dirichlet
    'ZETA_FUNCTIONS': 'ζ(s)',                 # Função zeta de Riemann
    'MODULAR_SYMBOLS': '{α, β}',              # Símbolos modulares
    'AUTOMORPHIC_FORMS': 'f(z)',              # Formas automórficas
    'REPRESENTATION_THEORY': 'ρ: G → GL(V)',  # Teoria de representações
    'LIE_ALGEBRAS': 'g',                      # Álgebras de Lie
    'QUANTUM_GROUPS': 'U_q(g)',               # Grupos quânticos
    'OPERAD_THEORY': 'O(n)',                  # Teoria de óperas
    'HIGHER_CATEGORIES': '∞-Cat',             # Categorias superiores
    'HOMOTOPY_THEORY': 'π_n(X)',              # Teoria de homotopia
    'STABLE_HOMOTOPY': 'π^S_*',               # Homotopia estável
    'SPECTRAL_SEQUENCES': 'E^{p,q}_r',        # Sequências espectrais
    'DERIVED_ALGEBRAIC': 'Spec R',            # Geometria algébrica derivada
    'PERFECTOID_SPACES': 'Spa(K, K⁺)',        # Espaços perfectoides
    'PRISMATIC_COHOMOLOGY': 'Δ',              # Cohomologia prismática
    'SYMPLECTIC_GEOMETRY': 'ω: V×V → ℝ',      # Geometria simplética
    'COMPLEX_GEOMETRY': '∂̄-operator',         # Geometria complexa
    'KAHLER_MANIFOLDS': 'g_{iȷ̄}',              # Variedades Kähler
    'CALABI-YAU': 'c₁ = 0',                   # Variedades Calabi-Yau
    'MIRROR_SYMMETRY': 'X ↔ X̂',                # Simetria espelho
    'STRING_THEORY': 'M^10 = M^4 × X^6',      # Teoria das cordas
    'M_THEORY': '11D Supergravity',           # Teoria M
    'F_THEORY': '12D Elliptic Fibration',     # Teoria F
    'TWISTOR_THEORY': 'ℙ𝕋',                    # Teoria twistor
    'NONCOMMUTATIVE': 'A_θ',                  # Geometria não-comutativa
    'QUANTUM_GRAVITY': 'ħG/c³ = ℓ_P²',        # Gravidade quântica
    'LOOP_QUANTUM': 'γ = ln(2)/π√3',          # Gravidade quântica em loop
    'SPIN_FOAMS': 'Γ → amplitudes',           # Espumas de spin
    'CAUSAL_SETS': '(C, ≺)',                  # Conjuntos causais
    'EMERGENT_GRAVITY': 'g_{μν} = ⟨T_{μν}⟩',  # Gravidade emergente
    'HOLOGRAPHIC': 'A = 4G_N S',              # Princípio holográfico
    'ADS_CFT': 'AdS_{d+1} ↔ CFT_d',           # Correspondência AdS/CFT
    'ENTANGLEMENT': 'S_A = -Tr(ρ_A log ρ_A)', # Entropia de emaranhamento
    'QUANTUM_ERROR': '[[n,k,d]]',             # Códigos quânticos
    'TOPOLOGICAL_Q': 'TQFT',                  # Computação quântica topológica
    'ANYONS': 'braiding statistics',          # Ánions
    'FRACTIONAL_Q': 'ν = p/q',                # Efeito Hall quântico fracionário
    'SPIN_LIQUIDS': 'QSL',                    # Líquidos de spin
    'SYMMETRY_PROTECTED': 'SPT phases',       # Fases SPT
    'TOPOLOGICAL_INSULATORS': 'Z₂ invariant', # Isolantes topológicos
    'WEGNER-WILSON': 'β(g) = μ ∂g/∂μ',        # Grupo de renormalização
    'CONFORMAL_FIELD': 'Δ, c',                # Teoria conformal de campos
    'INTEGRABLE_MODELS': 'R-matrix',          # Modelos integráveis
    'EXACTLY_SOLVABLE': 'Bethe ansatz',       # Modelos exatamente solúveis
    'STATISTICAL_MECHANICS': 'Z = Σ e^{-βE}', # Mecânica estatística
    'PHASE_TRANSITIONS': 'Tc',                # Transições de fase
    'CRITICAL_EXPONENTS': 'α, β, γ, δ, ν, η', # Expoentes críticos
    'RENORMALIZATION_GROUP': 'RG flow',       # Grupo de renormalização
    'UNIVERSALITY_CLASSES': 'Ising, XY, etc', # Classes de universalidade
    'SPONTANEOUS_SYMMETRY': 'φ → -φ',         # Quebra espontânea de simetria
    'GOLDSTONE_BOSONS': 'π(x)',               # Bósons de Goldstone
    'HIGGS_MECHANISM': 'A_μ → A_μ + ∂_μθ',    # Mecanismo de Higgs
    'ANOMALIES': '∂·j^5 ≠ 0',                 # Anomalias quânticas
    'INSTANTONS': '∫ F∧F = 8π²k',             # Instantons
    'MONOPOLES': '∫_S² F = 4πn',              # Monopolos magnéticos
    'VORTICES': '∮ A·dl = 2πn',               # Vórtices
    'DOMAIN_WALLS': 'φ(x) = v tanh(mx)',      # Paredes de domínio
    'COSMIC_STRINGS': 'μ ∼ v²',               # Cordas cósmicas
    'INFLATION': 'a(t) ∝ e^{Ht}',             # Inflação cósmica
    'DARK_ENERGY': 'Ω_Λ ≈ 0.69',              # Energia escura
    'DARK_MATTER': 'Ω_c ≈ 0.26',              # Matéria escura
    'BARYON_ASYMMETRY': 'η ≈ 6×10^{-10}',     # Assimetria bariônica
    'NEUTRINO_OSCILLATIONS': 'θ_{12}, θ_{23}, θ_{13}, δ_{CP}', # Oscilações de neutrinos
    'PROTON_DECAY': 'τ_p > 10^{34} anos',     # Decaimento do próton
    'NEUTRON_STARS': 'M ∼ 1.4 M_☉',           # Estrelas de nêutrons
    'BLACK_HOLES': 'R_S = 2GM/c²',            # Buracos negros
    'HAWKING_RADIATION': 'T_H = ħc³/8πGMk_B', # Radiação Hawking
    'INFORMATION_PARADOX': 'S_{BH} = A/4G_Nħ',# Paradoxo da informação
    'FIREWALL_PARADOX': 'AMPS',               # Paradoxo do firewall
    'ER=EPR': 'wormholes = entanglement',     # ER=EPR
    'QUANTUM_COMPLEMENTARITY': 'black hole complementarity', # Complementaridade quântica
    'HOLOGRAPHIC_NOISE': 'Δx ∼ √(ℓ_P L)',     # Ruído holográfico
    'NATURALNESS': 'hierarchy problem',       # Problema da hierarquia
    'FINE_TUNING': 'cosmological constant problem', # Problema da constante cosmológica
    'MULTIVERSE': 'landscape of vacua',       # Multiverso
    'ANTHROPIC_PRINCIPLE': 'selection effects', # Princípio antrópico
    'SIMULATION_HYPOTHESIS': 'Bostrom',       # Hipótese da simulação
    'MATHEMATICAL_UNIVERSE': 'Tegmark',       # Universo matemático
    'CONSCIOUSNESS_REALISM': 'Penrose-Hameroff', # Realismo da consciência
    'INTEGRATED_INFORMATION': 'Φ',            # Informação integrada
    'FREE_ENERGY_PRINCIPLE': 'Friston',       # Princípio da energia livre
    'PREDICTIVE_PROCESSING': 'brain as prediction machine', # Processamento preditivo
    'BAYESIAN_BRAIN': 'perception as inference', # Cérebro bayesiano
    'ACTIVE_INFERENCE': 'action as minimization of surprise', # Inferência ativa
    'MARKOV_BLANKETS': 'boundary of a system', # Cobertores de Markov
    'AUTOPOIETIC_SYSTEMS': 'self-producing systems', # Sistemas autopoiéticos
    'ENACTIVISM': 'cognition as embodied action', # Enativismo
    'EMBODIED_COGNITION': 'mind shaped by body', # Cognição incorporada
    'EXTENDED_MIND': 'mind beyond brain',     # Mente estendida
    'DISTRIBUTED_COGNITION': 'cognition across agents and artifacts', # Cognição distribuída
    'SITUATED_COGNITION': 'cognition in context', # Cognição situada
    'DYNAMICAL_SYSTEMS': 'ẋ = f(x)',          # Sistemas dinâmicos
    'ATTRACTORS': 'fixed points, limit cycles, strange attractors', # Atratores
    'BIFURCATIONS': 'pitchfork, Hopf, saddle-node', # Bifurcações
    'CHAOS_THEORY': 'butterfly effect',       # Teoria do caos
    'FRACTALS': 'self-similarity',            # Fractais
    'COMPLEXITY_THEORY': 'emergence, self-organization', # Teoria da complexidade
    'NETWORK_THEORY': 'graphs, small-world, scale-free', # Teoria de redes
    'INFORMATION_THEORY': 'Shannon entropy, mutual information', # Teoria da informação
    'ALGORITHMIC_INFORMATION': 'Kolmogorov complexity', # Informação algorítmica
    'COMPUTABILITY_THEORY': 'Turing machines, halting problem', # Teoria da computabilidade
    'COMPLEXITY_CLASSES': 'P, NP, BQP, etc',  # Classes de complexidade
    'QUANTUM_COMPUTATION': 'qubits, superposition, entanglement', # Computação quântica
    'QUANTUM_ALGORITHMS': "Shor's, Grover's", # Algoritmos quânticos
    'QUANTUM_ERROR_CORRECTION': 'stabilizer codes', # Correção de erro quântico
    'TOPOLOGICAL_QC': 'anyon braiding',       # Computação quântica topológica
    'QUANTUM_INFORMATION': 'density matrices, channels', # Informação quântica
    'QUANTUM_ENTANGLEMENT': 'Bell states, teleportation', # Emaranhamento quântico
    'QUANTUM_NONLOCALITY': "Bell's theorem",  # Não-localidade quântica
    'QUANTUM_CONTEXTUALITY': "Kochen-Specker", # Contextualidade quântica
    'QUANTUM_FOUNDATIONS': 'interpretations', # Fundamentos da quântica
    'DECOHERENCE': 'environment-induced superselection', # Decoerência
    'MEASUREMENT_PROBLEM': 'wave function collapse', # Problema da medição
    'MANY_WORLDS': 'Everett',                 # Muitos mundos
    'PILOT_WAVE': 'de Broglie-Bohm',          # Onda piloto
    'QBISM': 'quantum Bayesianism',           # QBism
    'RELATIONAL_QM': 'Rovelli',               # Quântica relacional
    'CONSCIOUSNESS_CAUSES': "von Neumann-Wigner", # Consciência causa colapso
    'OBJECTIVE_COLLAPSE': 'GRW, Penrose',     # Colapso objetivo
    'QUANTUM_THERMODYNAMICS': 'fluctuation theorems', # Termodinâmica quântica
    'NONEQUILIBRIUM_STAT_MECH': 'Jarzynski equality', # Mecânica estatística de não-equilíbrio
    'STOCHASTIC_THERMODYNAMICS': 'entropy production', # Termodinâmica estocástica
    'INFORMATION_THERMODYNAMICS': "Landauer's principle", # Termodinâmica da informação
    'MAXWELLS_DEMON': 'Szilard engine',       # Demônio de Maxwell
    'THERMODYNAMICS_OF_COMPUTATION': 'Bennett', # Termodinâmica da computação
    'QUANTUM_DISSIPATION': 'Caldeira-Leggett', # Dissipação quântica
    'OPEN_QUANTUM_SYSTEMS': 'Lindblad master equation', # Sistemas quânticos abertos
    'QUANTUM_OPTICS': 'Jaynes-Cummings model', # Óptica quântica
    'CAVITY_QED': 'strong coupling regime',   # QED de cavidade
    'TRAPPED_IONS': 'quantum gates',          # Íons presos
    'SUPERCONDUCTING_QUBITS': 'transmons, flux qubits', # Qubits supercondutores
    'TOPOLOGICAL_QUBITS': 'Majorana fermions',# Qubits topológicos
    'PHOTONIC_QC': 'linear optical QC',       # Computação quântica fotônica
    'NMR_QC': 'nuclear magnetic resonance QC', # Computação quântica por RMN
    'QUANTUM_DOTS': 'artificial atoms',       # Pontos quânticos
    'QUANTUM_HALL': 'anyon braiding platforms', # Efeito Hall quântico para QC
    'SPIN_QC': 'electron/nuclear spins',      # Computação quântica por spin
    'NEUTRAL_ATOMS': 'optical lattices, tweezer arrays', # Átomos neutros
    'MOLECULAR_QC': 'molecules as qubits',    # Computação quântica molecular
    'ADIABATIC_QC': 'quantum annealing',      # Computação quântica adiabática
    'QUANTUM_WALKS': 'quantum search algorithms', # Caminhadas quânticas
    'QUANTUM_MACHINE_LEARNING': 'quantum neural networks', # Aprendizado de máquina quântico
    'QUANTUM_CHEMISTRY': 'quantum simulation of molecules', # Química quântica
    'QUANTUM_MATERIALS': 'topological insulators, superconductors', # Materiais quânticos
    'QUANTUM_BIOLOGY': 'photosynthesis, magnetoreception', # Biologia quântica
    'QUANTUM_NEUROSCIENCE': 'quantum effects in brain', # Neurociência quântica
    'QUANTUM_COSMOLOGY': 'wave function of the universe', # Cosmologia quântica
    'QUANTUM_GRAVITY_LOOP': 'spin networks',  # Gravidade quântica em loop
    'STRING_THEORY_QUANTUM': 'string landscape', # Teoria das cordas quântica
    'HOLOGRAPHIC_DUALITY': 'gauge/gravity duality', # Dualidade holográfica
    'EMERGENT_SPACETIME': 'spacetime from entanglement', # Espaço-tempo emergente
    'QUANTUM_INFORMATION_IN_GRAVITY': 'black hole information paradox', # Informação quântica na gravidade
    'QUANTUM_FOUNDATIONS_OF_TIME': 'Page-Wootters, conditional probability', # Fundamentos quânticos do tempo
    'QUANTUM_CAUSAL_STRUCTURES': 'causal sets, quantum causal histories', # Estruturas causais quânticas
    'QUANTUM_REFERENCE_FRAMES': 'relational observables', # Referenciais quânticos
    'QUANTUM_CLOCKS': 'proper time from quantum systems', # Relógios quânticos
    'QUANTUM_THERMODYNAMICS_OF_GRAVITY': 'black hole thermodynamics', # Termodinâmica quântica da gravidade
    'QUANTUM_FLUCTUATIONS_IN_COSMOLOGY': 'inflationary perturbations', # Flutuações quânticas na cosmologia
    'QUANTUM_ORIGIN_OF_STRUCTURE': 'quantum to classical transition in cosmology', # Origem quântica da estrutura
    'QUANTUM_DECOOHERENCE_IN_COSMOLOGY': 'environment for cosmological perturbations', # Decoerência quântica na cosmologia
    'QUANTUM_INFORMATION_IN_COSMOLOGY': 'cosmological Bell tests', # Informação quântica na cosmologia
    'QUANTUM_GRAVITY_PHENOMENOLOGY': 'signatures of quantum gravity', # Fenomenologia da gravidade quântica
    'QUANTUM_FOAM': 'spacetime fluctuations at Planck scale', # Espuma quântica
    'NONCOMMUTATIVE_GEOMETRY_PHENOMENOLOGY': 'Lorentz violation, modified dispersion', # Fenomenologia da geometria não-comutativa
    'DOUBLY_SPECIAL_RELATIVITY': 'modified Lorentz transformations', # Relatividade duplamente especial
    'HORAVA-LIFSHITZ_GRAVITY': 'anisotropic scaling', # Gravidade de Horava-Lifshitz
    'ASYMPTOTIC_SAFETY': 'quantum gravity as QFT', # Segurança assintótica
    'CAUSAL_DYNAMICAL_TRIANGULATIONS': 'path integral for quantum gravity', # Triangulações dinâmicas causais
    'GROUP_FIELD_THEORY': 'quantum gravity as field theory on group manifold', # Teoria de campo de grupo
    'MATRIX_MODELS': 'M-theory in light-cone gauge', # Modelos de matriz
    'TENSOR_MODELS': 'higher-dimensional generalizations of matrix models', # Modelos tensoriais
    'SPIN_FOAM_MODELS': 'path integral for loop quantum gravity', # Modelos de espuma de spin
    'CANONICAL_LOOP_QUANTUM_GRAVITY': 'quantization of Ashtekar variables', # Gravidade quântica em loop canônica
    'LOOP_QUANTUM_COSMOLOGY': 'quantum cosmology from loop quantum gravity', # Cosmologia quântica em loop
    'STRING_GAS_COSMOLOGY': 'string theory early universe cosmology', # Cosmologia de gás de cordas
    'BRANE_WORLD_COSMOLOGY': 'brane inflation, ekpyrotic/cyclic models', # Cosmologia de mundo-brana
    'INFLATIONARY_COSMOLOGY': 'slow-roll inflation, eternal inflation', # Cosmologia inflacionária
    'BOUNCE_COSMOLOGIES': 'quantum bounce replacing big bang', # Cosmologias de quique
    'EMERGENT_UNIVERSE': 'universe from eternal static state', # Universo emergente
    'CONFORMAL_CYCLIC_COSMOLOGY': 'Penrose',  # Cosmologia cíclica conforme
    'MULTIVERSE_COSMOLOGY': 'bubble universes, landscape', # Cosmologia do multiverso
    'ANTHROPIC_COSMOLOGY': 'selection effects in multiverse', # Cosmologia antrópica
    'OBSERVER_EFFECTS_IN_COSMOLOGY': 'measurement problem in cosmology', # Efeitos do observador na cosmologia
    'QUANTUM_MEASUREMENTS_IN_COSMOLOGY': 'quantum state of the universe', # Medições quânticas na cosmologia
    'THERMODYNAMICS_OF_THE_UNIVERSE': 'cosmic entropy, arrow of time', # Termodinâmica do universo
    'INFORMATION_THEORY_IN_COSMOLOGY': 'cosmological information bounds', # Teoria da informação na cosmologia
    'COMPLEXITY_IN_COSMOLOGY': 'computational complexity of cosmological states', # Complexidade na cosmologia
    'HOLOGRAPHIC_COSMOLOGY': 'FRW/CFT correspondence', # Cosmologia holográfica
    'DUALITIES_IN_COSMOLOGY': 'dS/CFT, FRW/CFT', # Dualidades na cosmologia
    'OBSERVATIONAL_COSMOLOGY': 'CMB, LSS, gravitational waves', # Cosmologia observacional
    'COSMOLOGICAL_PARAMETERS': 'H0, Ωm, ΩΛ, σ8, ns, τ', # Parâmetros cosmológicos
    'COSMIC_MICROWAVE_BACKGROUND': 'anisotropies, polarization', # Radiação cósmica de fundo
    'LARGE_SCALE_STRUCTURE': 'galaxy clustering, BAO, redshift surveys', # Estrutura em grande escala
    'TYPE_IA_SUPERNOVAE': 'standard candles for cosmology', # Supernovas tipo Ia
    'GRAVITATIONAL_LENSING': 'weak, strong, microlensing', # Lentes gravitacionais
    'GRAVITATIONAL_WAVES': 'LIGO/Virgo, LISA, pulsar timing arrays', # Ondas gravitacionais
    '21CM_COSMOLOGY': 'neutral hydrogen from dark ages and reionization', # Cosmologia de 21cm
    'COSMIC_NEUTRINOS': 'CNB, neutrino mass constraints', # Neutrinos cósmicos
    'DARK_MATTER_SEARCHES': 'direct, indirect, collider', # Buscas por matéria escura
    'DARK_ENERGY_PROBES': 'SNe, BAO, WL, clusters', # Sondas de energia escura
    'INFLATIONARY_OBSERVABLES': 'tensor-to-scalar ratio, non-Gaussianity', # Observáveis inflacionários
    'PRIMORDIAL_BLACK_HOLES': 'dark matter candidates, early universe', # Buracos negros primordiais
    'COSMIC_STRINGS_OBSERVATIONS': 'CMB, gravitational waves, lensing', # Observações de cordas cósmicas
    'TESTS_OF_GRAVITY': 'solar system, binary pulsars, cosmology', # Testes da gravidade
    'VARYING_CONSTANTS': 'varying alpha, mu, me/mp', # Constantes variáveis
    'LORENTZ_VIOLATION': 'SME, astrophysical tests', # Violação de Lorentz
    'QUANTUM_GRAVITY_SIGNATURES': 'gamma-ray bursts, UHECR, gravitational waves', # Assinaturas da gravidade quântica
    'ASTROBIOLOGY': 'origin of life, habitable zones, biosignatures', # Astrobiologia
    'SETI': 'search for extraterrestrial intelligence', # SETI
    'FUTURE_COSMOLOGY_EXPERIMENTS': 'Euclid, LSST, SKA, CMB-S4, LISA', # Experimentos futuros de cosmologia
    'THEORY_OF_EVERYTHING': 'unification of all forces', # Teoria de tudo
    'FINAL_THEORY': 'ultimate laws of physics', # Teoria final
    'PHILOSOPHY_OF_PHYSICS': 'realism, empiricism, structuralism', # Filosofia da física
    'FOUNDATIONS_OF_PHYSICS': 'space, time, matter, causality', # Fundamentos da física
    'METAPHYSICS_OF_PHYSICS': 'nature of reality from physics', # Metafísica da física
    'PHYSICS_AND_MIND': 'consciousness in physical world', # Física e mente
    'PHYSICS_AND_MATHEMATICS': 'unreasonable effectiveness of mathematics', # Física e matemática
    'PHYSICS_AND_COMPUTATION': 'universe as computer', # Física e computação
    'PHYSICS_AND_INFORMATION': 'it from bit', # Física e informação
    'PHYSICS_AND_COMPLEXITY': 'emergence in physical systems', # Física e complexidade
    'PHYSICS_AND_SOCIETY': 'impact of physics on society', # Física e sociedade
    'FUTURE_OF_PHYSICS': 'next revolutions in physics', # Futuro da física
    'EDUCATION_IN_PHYSICS': 'teaching and learning physics', # Educação em física
    'HISTORY_OF_PHYSICS': 'development of physical theories', # História da física
    'PHYSICS_IN_ART': 'physics-inspired art', # Física na arte
    'PHYSICS_IN_LITERATURE': 'physics in fiction and non-fiction', # Física na literatura
    'PHYSICS_IN_MUSIC': 'physics of sound, music theory', # Física na música
    'PHYSICS_IN_SPORTS': 'physics of athletic performance', # Física no esporte
    'PHYSICS_IN_MEDICINE': 'medical physics, biophysics', # Física na medicina
    'PHYSICS_IN_ENGINEERING': 'applied physics, technology', # Física na engenharia
    'PHYSICS_IN_FINANCE': 'econophysics, quantitative finance', # Física nas finanças
    'PHYSICS_IN_BIOLOGY': 'biophysics, systems biology', # Física na biologia
    'PHYSICS_IN_CHEMISTRY': 'physical chemistry, chemical physics', # Física na química
    'PHYSICS_IN_MATERIALS_SCIENCE': 'condensed matter physics, nanotechnology', # Física na ciência dos materiais
    'PHYSICS_IN_ASTRONOMY': 'astrophysics, planetary science', # Física na astronomia
    'PHYSICS_IN_GEOLOGY': 'geophysics, seismology', # Física na geologia
    'PHYSICS_IN_CLIMATE_SCIENCE': 'climate physics, atmospheric physics', # Física na ciência do clima
    'PHYSICS_IN_ARCHAEOLOGY': 'archaeometry, dating techniques', # Física na arqueologia
    'PHYSICS_IN_FORENSICS': 'forensic physics, crime scene investigation', # Física na forense
    'PHYSICS_IN_FOOD_SCIENCE': 'food physics, culinary science', # Física na ciência dos alimentos
    'PHYSICS_IN_COSMETOLOGY': 'physics of cosmetics, skin care', # Física na cosmetologia
    'PHYSICS_IN_FASHION': 'physics of textiles, clothing design', # Física na moda
    'PHYSICS_IN_ARCHITECTURE': 'structural physics, building design', # Física na arquitetura
    'PHYSICS_IN_TRANSPORTATION': 'physics of vehicles, traffic flow', # Física no transporte
    'PHYSICS_IN_COMMUNICATION': 'physics of signals, information theory', # Física na comunicação
    'PHYSICS_IN_ENERGY': 'physics of energy production, storage', # Física na energia
    'PHYSICS_IN_ENVIRONMENT': 'environmental physics, pollution', # Física no ambiente
    'PHYSICS_IN_DEFENSE': 'military physics, weapons technology', # Física na defesa
    'PHYSICS_IN_SPACE_EXPLORATION': 'physics of space travel, colonization', # Física na exploração espacial
    'PHYSICS_IN_VIRTUAL_REALITY': 'physics engines, simulation', # Física na realidade virtual
    'PHYSICS_IN_GAMING': 'game physics, realistic graphics', # Física nos jogos
    'PHYSICS_IN_ANIMATION': 'physics-based animation, CGI', # Física na animação
    'PHYSICS_IN_ROBOTICS': 'physics of robots, control theory', # Física na robótica
    'PHYSICS_IN_AI': 'physics-inspired AI, neural networks', # Física na IA
    'PHYSICS_IN_BLOCKCHAIN': 'physics of consensus, cryptography', # Física no blockchain
    'PHYSICS_IN_QUANTUM_COMPUTING': 'physical implementation of qubits', # Física na computação quântica
    'PHYSICS_IN_NEUROSCIENCE': 'physics of the brain, neural dynamics', # Física na neurociência
    'PHYSICS_IN_PSYCHOLOGY': 'physics of behavior, decision making', # Física na psicologia
    'PHYSICS_IN_SOCIOLOGY': 'physics of social systems, networks', # Física na sociologia
    'PHYSICS_IN_ECONOMICS': 'econophysics, market dynamics', # Física na economia
    'PHYSICS_IN_POLITICAL_SCIENCE': 'physics of elections, power structures', # Física na ciência política
    'PHYSICS_IN_LINGUISTICS': 'physics of language, information theory', # Física na linguística
    'PHYSICS_IN_MUSICOLOGY': 'physics of music, acoustics', # Física na musicologia
    'PHYSICS_IN_ART_HISTORY': 'physics in art analysis, restoration', # Física na história da arte
    'PHYSICS_IN_PHILOSOPHY': 'physics and philosophy of science', # Física na filosofia
    'PHYSICS_IN_THEOLOGY': 'physics and religion, creation', # Física na teologia
    'PHYSICS_IN_MYSTICISM': 'physics and mysticism, consciousness', # Física no misticismo
    'PHYSICS_IN_ALCHEMY': 'physics and transformation, symbolism', # Física na alquimia
    'PHYSICS_IN_ASTROLOGY': 'physics and celestial influences', # Física na astrologia
    'PHYSICS_IN_MAGIC': 'physics and illusion, perception', # Física na mágica
    'PHYSICS_IN_SCIENCE_FICTION': 'physics in SF, speculative physics', # Física na ficção científica
    'PHYSICS_IN_FANTASY': 'physics in fantasy worlds, magic systems', # Física na fantasia
    'PHYSICS_IN_HORROR': 'physics in horror, supernatural phenomena', # Física no horror
    'PHYSICS_IN_COMEDY': 'physics in comedy, slapstick', # Física na comédia
    'PHYSICS_IN_DRAMA': 'physics in drama, human condition', # Física no drama
    'PHYSICS_IN_POETRY': 'physics in poetry, metaphor', # Física na poesia
    'PHYSICS_IN_PROSE': 'physics in prose, narrative', # Física na prosa
    'PHYSICS_IN_CINEMA': 'physics in film, special effects', # Física no cinema
    'PHYSICS_IN_THEATER': 'physics in theater, stage effects', # Física no teatro
    'PHYSICS_IN_OPERA': 'physics in opera, acoustics', # Física na ópera
    'PHYSICS_IN_DANCE': 'physics of dance, movement', # Física na dança
    'PHYSICS_IN_PAINTING': 'physics of painting, pigments', # Física na pintura
    'PHYSICS_IN_SCULPTURE': 'physics of sculpture, materials', # Física na escultura
    'PHYSICS_IN_PHOTOGRAPHY': 'physics of photography, light', # Física na fotografia
    'PHYSICS_IN_ARCHITECTURE': 'physics of architecture, structures', # Física na arquitetura
    'PHYSICS_IN_INDUSTRIAL_DESIGN': 'physics of design, ergonomics', # Física no design industrial
    'PHYSICS_IN_GRAPHIC_DESIGN': 'physics of graphics, color theory', # Física no design gráfico
    'PHYSICS_IN_WEB_DESIGN': 'physics of web, user experience', # Física no design web
    'PHYSICS_IN_GAME_DESIGN': 'physics in games, mechanics', # Física no design de jogos
    'PHYSICS_IN_PRODUCT_DESIGN': 'physics of products, functionality', # Física no design de produtos
    'PHYSICS_IN_FASHION_DESIGN': 'physics of fashion, textiles', # Física no design de moda
    'PHYSICS_IN_INTERIOR_DESIGN': 'physics of interiors, space', # Física no design de interiores
    'PHYSICS_IN_LANDSCAPE_DESIGN': 'physics of landscapes, ecology', # Física no design de paisagem
    'PHYSICS_IN_URBAN_DESIGN': 'physics of cities, planning', # Física no design urbano
    'PHYSICS_IN_TRANSPORT_DESIGN': 'physics of transport, vehicles', # Física no design de transporte
    'PHYSICS_IN_AEROSPACE_DESIGN': 'physics of aerospace, aircraft', # Física no design aeroespacial
    'PHYSICS_IN_MARINE_DESIGN': 'physics of marine, ships', # Física no design marítimo
    'PHYSICS_IN_AUTOMOTIVE_DESIGN': 'physics of automotive, cars', # Física no design automotivo
    'PHYSICS_IN_RAIL_DESIGN': 'physics of rail, trains', # Física no design ferroviário
    'PHYSICS_IN_BIKE_DESIGN': 'physics of bicycles, motorcycles', # Física no design de bicicletas
    'PHYSICS_IN_FOOTWEAR_DESIGN': 'physics of footwear, shoes', # Física no design de calçados
    'PHYSICS_IN_JEWELRY_DESIGN': 'physics of jewelry, gems', # Física no design de joias
    'PHYSICS_IN_WATCH_DESIGN': 'physics of watches, timekeeping', # Física no design de relógios
    'PHYSICS_IN_TOY_DESIGN': 'physics of toys, play', # Física no design de brinquedos
    'PHYSICS_IN_FURNITURE_DESIGN': 'physics of furniture, comfort', # Física no design de móveis
    'PHYSICS_IN_KITCHEN_DESIGN': 'physics of kitchens, appliances', # Física no design de cozinhas
    'PHYSICS_IN_BATHROOM_DESIGN': 'physics of bathrooms, plumbing', # Física no design de banheiros
    'PHYSICS_IN_BEDROOM_DESIGN': 'physics of bedrooms, sleep', # Física no design de quartos
    'PHYSICS_IN_LIVING_ROOM_DESIGN': 'physics of living rooms, socializing', # Física no design de salas de estar
    'PHYSICS_IN_DINING_ROOM_DESIGN': 'physics of dining rooms, eating', # Física no design de salas de jantar
    'PHYSICS_IN_OFFICE_DESIGN': 'physics of offices, work', # Física no design de escritórios
    'PHYSICS_IN_STUDIO_DESIGN': 'physics of studios, creativity', # Física no design de estúdios
    'PHYSICS_IN_GARAGE_DESIGN': 'physics of garages, storage', # Física no design de garagens
    'PHYSICS_IN_BASEMENT_DESIGN': 'physics of basements, utility', # Física no design de porões
    'PHYSICS_IN_ATTIC_DESIGN': 'physics of attics, insulation', # Física no design de sótãos
    'PHYSICS_IN_GARDEN_DESIGN': 'physics of gardens, growth', # Física no design de jardins
    'PHYSICS_IN_PATIO_DESIGN': 'physics of patios, outdoor living', # Física no design de pátios
    'PHYSICS_IN_POOL_DESIGN': 'physics of pools, water', # Física no design de piscinas
    'PHYSICS_IN_SPA_DESIGN': 'physics of spas, relaxation', # Física no design de spas
    'PHYSICS_IN_GYM_DESIGN': 'physics of gyms, exercise', # Física no design de academias
    'PHYSICS_IN_SAUNA_DESIGN': 'physics of saunas, heat', # Física no design de saunas
    'PHYSICS_IN_STEAM_ROOM_DESIGN': 'physics of steam rooms, humidity', # Física no design de banhos a vapor
    'PHYSICS_IN_MEDITATION_ROOM_DESIGN': 'physics of meditation rooms, peace', # Física no design de salas de meditação
    'PHYSICS_IN_YOGA_ROOM_DESIGN': 'physics of yoga rooms, flexibility', # Física no design de salas de yoga
    'PHYSICS_IN_DANCE_ROOM_DESIGN': 'physics of dance rooms, movement', # Física no design de salas de dança
    'PHYSICS_IN_MUSIC_ROOM_DESIGN': 'physics of music rooms, acoustics', # Física no design de salas de música
    'PHYSICS_IN_ART_ROOM_DESIGN': 'physics of art rooms, creativity', # Física no design de salas de arte
    'PHYSICS_IN_CRAFT_ROOM_DESIGN': 'physics of craft rooms, making', # Física no design de salas de artesanato
    'PHYSICS_IN_SEWING_ROOM_DESIGN': 'physics of sewing rooms, textiles', # Física no design de salas de costura
    'PHYSICS_IN_WOODWORKING_ROOM_DESIGN': 'physics of woodworking rooms, wood', # Física no design de salas de marcenaria
    'PHYSICS_IN_METALWORKING_ROOM_DESIGN': 'physics of metalworking rooms, metal', # Física no design de salas de metalurgia
    'PHYSICS_IN_POTTERY_ROOM_DESIGN': 'physics of pottery rooms, clay', # Física no design de salas de cerâmica
    'PHYSICS_IN_GLASSBLOWING_ROOM_DESIGN': 'physics of glassblowing rooms, glass', # Física no design de salas de sopro de vidro
    'PHYSICS_IN_PRINTING_ROOM_DESIGN': 'physics of printing rooms, ink', # Física no design de salas de impressão
    'PHYSICS_IN_DARKROOM_DESIGN': 'physics of darkrooms, photography', # Física no design de câmaras escuras
    'PHYSICS_IN_LABORATORY_DESIGN': 'physics of laboratories, experimentation', # Física no design de laboratórios
    'PHYSICS_IN_WORKSHOP_DESIGN': 'physics of workshops, repair', # Física no design de oficinas
    'PHYSICS_IN_STORAGE_ROOM_DESIGN': 'physics of storage rooms, organization', # Física no design de salas de armazenamento
    'PHYSICS_IN_CLOSET_DESIGN': 'physics of closets, clothing', # Física no design de armários
    'PHYSICS_IN_PANTRY_DESIGN': 'physics of pantries, food', # Física no design de despensas
    'PHYSICS_IN_CELLAR_DESIGN': 'physics of cellars, wine', # Física no design de adegas
    'PHYSICS_IN_VAULT_DESIGN': 'physics of vaults, security', # Física no design de cofres
    'PHYSICS_IN_SAFE_ROOM_DESIGN': 'physics of safe rooms, protection', # Física no design de salas seguras
    'PHYSICS_IN_BUNKER_DESIGN': 'physics of bunkers, survival', # Física no design de bunkers
    'PHYSICS_IN_SHELTER_DESIGN': 'physics of shelters, emergency', # Física no design de abrigos
    'PHYSICS_IN_TENT_DESIGN': 'physics of tents, camping', # Física no design de tendas
    'PHYSICS_IN_RV_DESIGN': 'physics of RVs, mobile living', # Física no design de veículos recreativos
    'PHYSICS_IN_BOAT_DESIGN': 'physics of boats, sailing', # Física no design de barcos
    'PHYSICS_IN_PLANE_DESIGN': 'physics of planes, flying', # Física no design de aviões
    'PHYSICS_IN_CAR_DESIGN': 'physics of cars, driving', # Física no design de carros
    'PHYSICS_IN_BIKE_DESIGN': 'physics of bikes, cycling', # Física no design de bicicletas
    'PHYSICS_IN_SCOOTER_DESIGN': 'physics of scooters, mobility', # Física no design de scooters
    'PHYSICS_IN_SKATEBOARD_DESIGN': 'physics of skateboards, tricks', # Física no design de skates
    'PHYSICS_IN_SKI_DESIGN': 'physics of skis, snow', # Física no design de esquis
    'PHYSICS_IN_SNOWBOARD_DESIGN': 'physics of snowboards, powder', # Física no design de snowboards
    'PHYSICS_IN_SURFBOARD_DESIGN': 'physics of surfboards, waves', # Física no design de pranchas de surf
    'PHYSICS_IN_KITE_DESIGN': 'physics of kites, wind', # Física no design de pipas
    'PHYSICS_IN_PARAGLIDER_DESIGN': 'physics of paragliders, air', # Física no design de parapentes
    'PHYSICS_IN_HANG_GLIDER_DESIGN': 'physics of hang gliders, soaring', # Física no design de asas delta
    'PHYSICS_IN_HOT_AIR_BALLOON_DESIGN': 'physics of hot air balloons, buoyancy', # Física no design de balões de ar quente
}
```

---
**TODO: Completing the truncated list of constants (ended at PHYSICS_IN_ZE).**
