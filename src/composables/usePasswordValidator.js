import { ref, computed } from "vue";

export function usePasswordValidator() {
  const password = ref("");

  const passwordStrength = computed(() => {
    const p = password.value;
    if (p.length === 0)
      return { score: 0, text: "", color: "#ccc", width: "0%" };

    let score = 0;
    let checks = 0;

    if (p.length >= 8) {
      score += 20;
      checks++;
    }
    if (/[a-z]/.test(p)) {
      score += 20;
      checks++;
    }
    if (/[A-Z]/.test(p)) {
      score += 20;
      checks++;
    }
    if (/[0-9]/.test(p)) {
      score += 20;
      checks++;
    }
    if (/[^a-zA-Z0-9]/.test(p)) {
      score += 20;
      checks++;
    }

    score = Math.min(score, 100);

    let text = "Sangat Lemah";
    let color = "#dc3545";

    if (checks === 5) {
      text = "Sangat Kuat";
      color = "#28a745";
    } else if (checks === 4) {
      text = "Kuat";
      color = "#8fce00";
    } else if (checks === 3) {
      text = "Kurang Kuat";
      color = "#ffc107";
    } else if (checks === 2) {
      text = "Lemah";
      color = "#fd7e14";
    }

    return { score, text, color, width: `${score}%` };
  });

  function validatePassword(passwordToValidate) {
    if (passwordToValidate.length < 8)
      return "Kata sandi harus memiliki minimal 8 karakter.";
    if (!/[a-z]/.test(passwordToValidate))
      return "Kata sandi harus mengandung setidaknya satu huruf kecil.";
    if (!/[A-Z]/.test(passwordToValidate))
      return "Kata sandi harus mengandung setidaknya satu huruf besar.";
    if (!/[0-9]/.test(passwordToValidate))
      return "Kata sandi harus mengandung setidaknya satu angka.";
    if (!/[^a-zA-Z0-9]/.test(passwordToValidate))
      return "Kata sandi harus mengandung setidaknya satu karakter spesial.";
    return null;
  }

  return { password, passwordStrength, validatePassword };
}
