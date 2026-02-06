//! Fundamental Constants of Avalon
//! Scaled by 1000 for fixed-point arithmetic.

/// Estado fundamental 7.0
pub const GROUND_STATE_7: i64 = 7000;

/// Minutos de santuário
pub const SANCTUARY_TIME: i64 = 144;

/// Minutos máximos por gesto
pub const ATOMIC_GESTURE_MAX: i64 = 5;

/// Δ mínimo significativo (0.33)
pub const QUANTUM_LEAP_THRESHOLD: i64 = 330;

/// % para exclusão automática (0.95)
pub const EXCLUSION_THRESHOLD: i64 = 950;

/// Hz de ressonância (144.963)
pub const FIELD_COHERENCE: i64 = 144963;

/// π×10 ≈ 31.4 (Bitcoin 31.x) (31.415)
pub const SATOSHI_FREQUENCY: i64 = 31415;

/// Å (parâmetro de rede do diamante) (3.567)
pub const DIAMOND_LATTICE_CONSTANT: i64 = 3567;

/// Proporção áurea (1.618)
pub const GOLDEN_RATIO_FIELD: i64 = 1618;

/// π (completude) (3.141)
pub const PI_CIRCULARITY: i64 = 3141;

/// e (crescimento orgânico) (2.718)
pub const EULER_IDENTITY: i64 = 2718;

/// BTC/unidade atômica (1e-8) -> In BTC units, it's 0.00000001
/// Scaled by 1000 it would be 0.00001 which is still small.
/// For satoshis, we usually use 1.
pub const SATOSHI_SATOSHI: i64 = 1;
