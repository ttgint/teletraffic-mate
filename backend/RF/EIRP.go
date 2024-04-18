package RF;

func Eirp(TransmittedPower, CombinerLoss, FeederLoss, AntennaGain float64) float64 {
    return TransmittedPower - CombinerLoss - FeederLoss + AntennaGain
}

func TransmittedPower(Eirp, CombinerLoss, FeederLoss, AntennaGain float64) float64 {
    return Eirp + CombinerLoss + FeederLoss - AntennaGain
}

func CombinerLoss(Eirp, TransmittedPower, FeederLoss, AntennaGain float64) float64 {
    return TransmittedPower - FeederLoss + AntennaGain - Eirp
}

func FeederLoss(Eirp, TransmittedPower, CombinerLoss, AntennaGain float64) float64 {
    return TransmittedPower - CombinerLoss + AntennaGain - Eirp
}

func AntennaGain(Eirp, TransmittedPower, CombinerLoss, FeederLoss float64) float64 {
    return Eirp - TransmittedPower + CombinerLoss + FeederLoss
}
