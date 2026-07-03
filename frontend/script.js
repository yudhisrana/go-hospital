const API_BASE = "http://localhost:8080";

const state = {
  prescriptionCount: 0,
  currentFilter: "all",
  patients: [],
  allVisits: [],
  filteredVisits: [],
  visits: [],
  medicineSuggestions: ["Paracetamol", "Amoxicillin", "Ibuprofen"],
};

document.addEventListener("DOMContentLoaded", () => {
  initTabs();
  initPickers();
  addPrescriptionField();
  refreshDashboard();
  loadAllVisits();
  loadPatientsForPicker();
  loadVisitPickers();
  renderMedicineSuggestions();
});

function initTabs() {
  const buttons = document.querySelectorAll(".tab-button");
  const tabs = document.querySelectorAll(".tab-content");

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      const target = button.dataset.tab;

      buttons.forEach((b) => b.classList.remove("active"));
      tabs.forEach((tab) => tab.classList.remove("active"));

      button.classList.add("active");
      document.getElementById(target).classList.add("active");
    });
  });
}

function setActiveFilter(filterName) {
  state.currentFilter = filterName;
  document.querySelectorAll(".filter-btn").forEach((button) => {
    button.classList.toggle("active", button.dataset.filter === filterName);
  });
}

function initPickers() {
  const patientPicker = document.getElementById("patientPicker");
  if (patientPicker) {
    patientPicker.addEventListener("change", (event) => {
      const id = event.target.value;
      const patient = state.patients.find((p) => p.id === id);
      if (!patient) return;
      selectPatient(patient.id, patient.name);
    });
  }

  const examineVisitPicker = document.getElementById("examineVisitPicker");
  if (examineVisitPicker) {
    examineVisitPicker.addEventListener("change", (event) => {
      const id = event.target.value;
      if (id) {
        document.getElementById("visitId").value = id;
      }
      renderVisitSummary("visitId", "examineVisitSummary");
    });
  }

  const examineVisitId = document.getElementById("visitId");
  if (examineVisitId) {
    examineVisitId.addEventListener("input", () => {
      syncPickerSelection("examineVisitPicker", examineVisitId.value);
      renderVisitSummary("visitId", "examineVisitSummary");
    });
  }

  const dispenseVisitPicker = document.getElementById("dispenseVisitPicker");
  if (dispenseVisitPicker) {
    dispenseVisitPicker.addEventListener("change", (event) => {
      const id = event.target.value;
      if (id) {
        document.getElementById("dispenseVisitId").value = id;
      }
      renderVisitSummary("dispenseVisitId", "dispenseVisitSummary");
    });
  }

  const dispenseVisitId = document.getElementById("dispenseVisitId");
  if (dispenseVisitId) {
    dispenseVisitId.addEventListener("input", () => {
      syncPickerSelection("dispenseVisitPicker", dispenseVisitId.value);
      renderVisitSummary("dispenseVisitId", "dispenseVisitSummary");
    });
  }
}

function goToTab(tabId) {
  const button = document.querySelector(`.tab-button[data-tab="${tabId}"]`);
  if (button) {
    button.click();
  }
}

function renderMedicineSuggestions() {
  const datalist = document.getElementById("medicine-suggestions");
  if (!datalist) return;

  const unique = Array.from(
    new Set(state.medicineSuggestions.filter(Boolean)),
  ).sort();
  datalist.innerHTML = unique
    .map((name) => `<option value="${name}"></option>`)
    .join("");
}

function updateMedicineSuggestionsFromVisits(visits) {
  const names = [];
  (visits || []).forEach((visit) => {
    (visit.prescriptions || []).forEach((item) => {
      if (item?.medicine) names.push(item.medicine);
    });
  });

  if (names.length) {
    state.medicineSuggestions = [...state.medicineSuggestions, ...names];
    renderMedicineSuggestions();
  }
}

function getPatientDisplayName(patientId) {
  const patient = state.patients.find((item) => item.id === patientId);
  if (!patient) return patientId;
  return `${patient.name} (${patient.nik})`;
}

function getPatientDetails(patientId) {
  const patient = state.patients.find((item) => item.id === patientId);
  if (!patient) {
    return {
      name: patientId,
      nik: "-",
      gender: "-",
      age: "-",
      phone_number: "-",
      address: "-",
    };
  }

  return patient;
}

function getVisitById(visitId) {
  const visits = state.allVisits || state.visits || [];
  return visits.find((visit) => visit.id === visitId) || null;
}

function syncPickerSelection(pickerId, visitId) {
  const picker = document.getElementById(pickerId);
  if (!picker) return;
  picker.value = visitId || "";
}

function renderVisitSummary(inputId, summaryId) {
  const input = document.getElementById(inputId);
  const summary = document.getElementById(summaryId);
  if (!input || !summary) return;

  const visitId = input.value.trim();
  if (!visitId) {
    summary.className = "visit-summary-box";
    summary.innerHTML =
      "Pilih kunjungan untuk melihat detail pasien dan kunjungan.";
    return;
  }

  const visit = getVisitById(visitId);
  if (!visit) {
    summary.className = "visit-summary-box warning";
    summary.innerHTML = `Kunjungan <strong>${visitId}</strong> belum ditemukan di daftar.`;
    return;
  }

  const patient = getPatientDetails(visit.patient_id);
  const statusLabel = visit.status ? visit.status.replaceAll("_", " ") : "-";
  const prescriptions = (visit.prescriptions || [])
    .map((item) => `${item.medicine} • ${item.dosage} • Qty ${item.quantity}`)
    .join("<br />");

  summary.className = "visit-summary-box has-data";
  summary.innerHTML = `
    <div class="visit-summary-title">Detail Pasien dan Kunjungan</div>
    <div class="visit-summary-grid">
      <div><span>Nama</span><strong>${patient.name || "-"}</strong></div>
      <div><span>NIK</span><strong>${patient.nik || "-"}</strong></div>
      <div><span>Gender</span><strong>${patient.gender || "-"}</strong></div>
      <div><span>Umur</span><strong>${patient.age || "-"}</strong></div>
      <div><span>Telepon</span><strong>${patient.phone_number || patient.phone || "-"}</strong></div>
      <div><span>Status Visit</span><strong>${statusLabel}</strong></div>
    </div>
    <div class="visit-summary-notes">
      <div><span>Symptoms</span><p>${visit.symptoms || "-"}</p></div>
      <div><span>Diagnosis</span><p>${visit.diagnosis || "-"}</p></div>
      <div><span>Prescription</span><p>${prescriptions || "-"}</p></div>
    </div>
  `;
}

function renderPatientPickerOptions(patients) {
  const picker = document.getElementById("patientPicker");
  if (!picker) return;

  if (!patients.length) {
    picker.innerHTML = '<option value="">Belum ada data pasien</option>';
    return;
  }

  picker.innerHTML =
    '<option value="">Pilih pasien...</option>' +
    patients
      .map((p) => `<option value="${p.id}">${p.name} - ${p.nik}</option>`)
      .join("");
}

function renderVisitPickerOptions() {
  const examinePicker = document.getElementById("examineVisitPicker");
  const dispensePicker = document.getElementById("dispenseVisitPicker");
  const visits = state.visits || [];

  const registered = visits.filter((visit) => visit.status === "registered");
  const waitingPharmacy = visits.filter(
    (visit) => visit.status === "waiting_pharmacy",
  );

  if (examinePicker) {
    examinePicker.innerHTML =
      '<option value="">Pilih kunjungan untuk diperiksa...</option>' +
      registered
        .map(
          (v) =>
            `<option value="${v.id}">Queue #${v.queue_number} - ${v.patient_id}</option>`,
        )
        .join("");
    if (registered.length === 0) {
      examinePicker.innerHTML =
        '<option value="">Tidak ada kunjungan status registered</option>';
    }
  }

  if (dispensePicker) {
    dispensePicker.innerHTML =
      '<option value="">Pilih kunjungan untuk obat...</option>' +
      waitingPharmacy
        .map(
          (v) =>
            `<option value="${v.id}">Queue #${v.queue_number} - ${v.patient_id}</option>`,
        )
        .join("");
    if (waitingPharmacy.length === 0) {
      dispensePicker.innerHTML =
        '<option value="">Tidak ada kunjungan status waiting_pharmacy</option>';
    }
  }
}

async function loadPatientsForPicker() {
  try {
    const result = await apiFetch("/patients");
    state.patients = result.data || [];
    renderPatientPickerOptions(state.patients);
    renderVisitSummary("visitId", "examineVisitSummary");
    renderVisitSummary("dispenseVisitId", "dispenseVisitSummary");
  } catch {
    renderPatientPickerOptions([]);
  }
}

async function loadVisitPickers() {
  try {
    const result = await apiFetch("/visits");
    state.allVisits = result.data || [];
    state.visits = state.allVisits;
    updateMedicineSuggestionsFromVisits(state.allVisits);
    renderVisitPickerOptions();
    renderVisitSummary("visitId", "examineVisitSummary");
    renderVisitSummary("dispenseVisitId", "dispenseVisitSummary");
  } catch {
    state.allVisits = [];
    state.visits = [];
    renderVisitPickerOptions();
  }
}

function startExamineFromVisit(visitId) {
  document.getElementById("visitId").value = visitId;
  const picker = document.getElementById("examineVisitPicker");
  if (picker) picker.value = visitId;
  goToTab("examine");
}

function startDispenseFromVisit(visitId) {
  document.getElementById("dispenseVisitId").value = visitId;
  const picker = document.getElementById("dispenseVisitPicker");
  if (picker) picker.value = visitId;
  goToTab("dispense");
}

function animateCounter(id, value) {
  const el = document.getElementById(id);
  if (!el) return;

  const from = Number(el.textContent) || 0;
  const to = Number(value) || 0;
  const duration = 360;
  const start = performance.now();

  function frame(now) {
    const progress = Math.min((now - start) / duration, 1);
    const current = Math.round(from + (to - from) * progress);
    el.textContent = current;
    if (progress < 1) requestAnimationFrame(frame);
  }

  requestAnimationFrame(frame);
}

function formatResponse(data) {
  return JSON.stringify(data, null, 2);
}

function normalizeBirthDate(value) {
  const raw = (value || "").trim();
  if (!raw) return "";

  // Accept YYYY-MM-DD from date input and convert to RFC3339 expected by backend.
  if (/^\d{4}-\d{2}-\d{2}$/.test(raw)) {
    return `${raw}T00:00:00Z`;
  }

  // Keep valid ISO datetime as-is.
  const parsed = new Date(raw);
  if (!Number.isNaN(parsed.getTime())) {
    return parsed.toISOString();
  }

  return "";
}

function calculateAgeFromDateInput(dateValue) {
  if (!dateValue) return 0;

  const birth = new Date(`${dateValue}T00:00:00Z`);
  if (Number.isNaN(birth.getTime())) return 0;

  const today = new Date();
  let age = today.getUTCFullYear() - birth.getUTCFullYear();
  const monthDiff = today.getUTCMonth() - birth.getUTCMonth();
  const dayDiff = today.getUTCDate() - birth.getUTCDate();

  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age -= 1;
  }

  return age > 0 ? age : 0;
}

async function apiFetch(path, options = {}) {
  const res = await fetch(`${API_BASE}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
    },
    ...options,
  });

  const text = await res.text();
  let payload;
  try {
    payload = text ? JSON.parse(text) : null;
  } catch {
    payload = { raw: text };
  }

  if (!res.ok) {
    const message = payload?.message || payload?.error || `HTTP ${res.status}`;
    throw new Error(message);
  }

  return payload;
}

function setBoxContent(selector, message, type = "success") {
  const box = document.querySelector(selector);
  if (!box) return;
  box.className = `response-box ${type}`;
  box.textContent =
    typeof message === "string" ? message : formatResponse(message);
}

function clearBox(selector) {
  const box = document.querySelector(selector);
  if (!box) return;
  box.className = "response-box";
  box.textContent = "";
}

async function registerVisit() {
  const patientId = document.getElementById("patientId").value.trim();
  const symptoms = document.getElementById("symptoms").value.trim();

  if (!patientId || !symptoms) {
    setBoxContent(
      "#registerResponse",
      "ID pasien dan gejala wajib diisi.",
      "error",
    );
    return;
  }

  try {
    const result = await apiFetch("/visits/register", {
      method: "POST",
      body: JSON.stringify({ patient_id: patientId, symptoms }),
    });
    setBoxContent("#registerResponse", result, "success");
    document.getElementById("registerForm").reset();
    await refreshDashboard();
    await loadAllVisits();
    await loadVisitPickers();
  } catch (error) {
    setBoxContent("#registerResponse", error.message, "error");
  }
}

// Toggle create patient box
function toggleCreatePatient() {
  const box = document.getElementById("createPatientBox");
  if (!box) return;
  box.classList.toggle("is-hidden");
}

// Search patient by NIK or name (client-side filtering)
async function searchPatient() {
  const q = document.getElementById("patientQuery").value.trim().toLowerCase();
  const resultBox = document.getElementById("patientSearchResult");
  resultBox.textContent = "";
  if (!q) {
    resultBox.textContent = "Masukkan NIK atau nama untuk mencari.";
    return;
  }

  try {
    const res = await apiFetch("/patients");
    const patients = res.data || [];
    const matched = patients.filter(
      (p) =>
        (p.nik && p.nik.toLowerCase().includes(q)) ||
        (p.name && p.name.toLowerCase().includes(q)),
    );
    if (matched.length === 0) {
      resultBox.textContent = "Pasien tidak ditemukan.";
      return;
    }
    // if only one match, select it
    if (matched.length === 1) {
      document.getElementById("patientId").value = matched[0].id;
      resultBox.textContent = `Pasien dipilih: ${matched[0].name} (ID: ${matched[0].id})`;
      return;
    }

    // multiple matches: show clickable list
    resultBox.innerHTML = matched
      .map(
        (p) =>
          `<div style="margin-bottom:6px;"><button class="btn btn-secondary small-btn" onclick="selectPatient('${p.id}','${p.name.replace(/'/g, "\'")}')">Pilih</button> &nbsp; ${p.name} — ${p.nik} — <small>${p.id}</small></div>`,
      )
      .join("");
  } catch (err) {
    resultBox.textContent = "Gagal mencari pasien: " + err.message;
  }
}

function selectPatient(id, name) {
  document.getElementById("patientId").value = id;
  document.getElementById("patientSearchResult").textContent =
    `Pasien dipilih: ${name} (ID: ${id})`;
}

// Create a new patient via API and auto-select
async function createPatient() {
  const nik = document.getElementById("new_nik").value.trim();
  const name = document.getElementById("new_name").value.trim();
  const gender = document.getElementById("new_gender").value.trim();
  const birthDateInput = document.getElementById("new_birthdate").value;
  const age = calculateAgeFromDateInput(birthDateInput);
  const birth_date = normalizeBirthDate(birthDateInput);
  const address = document.getElementById("new_address").value.trim();
  const phone_number = document.getElementById("new_phone").value.trim();

  if (
    !nik ||
    !name ||
    !age ||
    !gender ||
    !birth_date ||
    !address ||
    !phone_number
  ) {
    setBoxContent(
      "#registerResponse",
      "Semua field pasien baru harus diisi dengan format valid.",
      "error",
    );
    return;
  }

  try {
    const res = await apiFetch("/patients", {
      method: "POST",
      body: JSON.stringify({
        nik,
        name,
        age,
        gender,
        birth_date,
        address,
        phone_number,
      }),
    });
    const patient = res.data || res;
    // set patientId and close create box
    if (patient && patient.id) {
      document.getElementById("patientId").value = patient.id;
      setBoxContent(
        "#registerResponse",
        `Pasien dibuat dan dipilih: ${patient.name} (ID: ${patient.id})`,
        "success",
      );
      document.getElementById("createPatientBox").classList.add("is-hidden");
      document.getElementById("patientSearchResult").textContent = "";
      await loadPatientsForPicker();
    } else {
      setBoxContent(
        "#registerResponse",
        "Pasien dibuat, namun response tidak mengandung ID.",
        "error",
      );
    }
  } catch (err) {
    setBoxContent(
      "#registerResponse",
      "Gagal membuat pasien: " + err.message,
      "error",
    );
  }
}

function addPrescriptionField() {
  state.prescriptionCount += 1;
  const wrapper = document.getElementById("prescriptionFields");
  const id = `prescription-${state.prescriptionCount}`;

  const row = document.createElement("div");
  row.className = "prescription-item";
  row.id = id;
  row.innerHTML = `
        <input type="text" class="medicine-name" list="medicine-suggestions" placeholder="Medicine" />
        <input type="text" class="medicine-dosage" placeholder="Dosage" />
        <input type="number" class="medicine-qty" placeholder="Qty" min="1" />
        <button type="button" class="btn btn-secondary small-btn" onclick="removePrescriptionField('${id}')">Remove</button>
    `;
  wrapper.appendChild(row);
}

function removePrescriptionField(id) {
  const el = document.getElementById(id);
  if (el) el.remove();
}

async function examinePatient() {
  const visitId = document.getElementById("visitId").value.trim();
  const diagnosis = document.getElementById("diagnosis").value.trim();
  const rows = document.querySelectorAll(
    "#prescriptionFields .prescription-item",
  );

  if (!visitId || !diagnosis) {
    setBoxContent(
      "#examineResponse",
      "Visit ID dan diagnosis wajib diisi.",
      "error",
    );
    return;
  }

  const prescriptions = [];
  rows.forEach((row) => {
    const medicine = row.querySelector(".medicine-name").value.trim();
    const dosage = row.querySelector(".medicine-dosage").value.trim();
    const quantity = Number(row.querySelector(".medicine-qty").value);

    if (medicine && dosage && quantity > 0) {
      prescriptions.push({ medicine, dosage, quantity });
    }
  });

  if (prescriptions.length === 0) {
    setBoxContent("#examineResponse", "Minimal 1 resep harus diisi.", "error");
    return;
  }

  try {
    const result = await apiFetch(`/visits/${visitId}/examine`, {
      method: "PATCH",
      body: JSON.stringify({ diagnosis, prescriptions }),
    });
    setBoxContent("#examineResponse", result, "success");
    document.getElementById("examineForm").reset();
    document.getElementById("prescriptionFields").innerHTML = "";
    state.prescriptionCount = 0;
    addPrescriptionField();
    await refreshDashboard();
    await loadAllVisits();
    await loadVisitPickers();
  } catch (error) {
    setBoxContent("#examineResponse", error.message, "error");
  }
}

async function dispenseMedicine() {
  const visitId = document.getElementById("dispenseVisitId").value.trim();

  if (!visitId) {
    setBoxContent("#dispenseResponse", "Visit ID wajib diisi.", "error");
    return;
  }

  try {
    const result = await apiFetch(`/visits/${visitId}/dispense`, {
      method: "PATCH",
    });
    setBoxContent("#dispenseResponse", result, "success");
    document.getElementById("dispenseForm").reset();
    await refreshDashboard();
    await loadAllVisits();
    await loadVisitPickers();
  } catch (error) {
    setBoxContent("#dispenseResponse", error.message, "error");
  }
}

async function loadAllVisits() {
  try {
    const result = await apiFetch("/visits");
    state.allVisits = result.data || [];
    setActiveFilter("all");
    renderVisits(state.allVisits);
    renderVisitSummary("visitId", "examineVisitSummary");
    renderVisitSummary("dispenseVisitId", "dispenseVisitSummary");
  } catch (error) {
    const list = document.getElementById("visitsList");
    list.innerHTML = `<div class="response-box error">${error.message}</div>`;
  }
}

async function loadVisitsByStatus(status) {
  try {
    const result = await apiFetch(
      `/visits/status/${encodeURIComponent(status)}`,
    );
    setActiveFilter(status);
    renderVisits(result.data || []);
    renderVisitSummary("visitId", "examineVisitSummary");
    renderVisitSummary("dispenseVisitId", "dispenseVisitSummary");
  } catch (error) {
    const list = document.getElementById("visitsList");
    list.innerHTML = `<div class="response-box error">${error.message}</div>`;
  }
}

function renderVisits(visits) {
  const list = document.getElementById("visitsList");
  state.filteredVisits = visits || [];
  state.visits = state.allVisits.length
    ? state.allVisits
    : state.filteredVisits;
  updateMedicineSuggestionsFromVisits(
    state.allVisits.length ? state.allVisits : state.filteredVisits,
  );
  renderVisitPickerOptions();

  if (!visits || visits.length === 0) {
    list.innerHTML = '<div class="visit-card">Belum ada visit.</div>';
    return;
  }

  list.innerHTML = visits
    .map((visit) => {
      const prescriptions = (visit.prescriptions || [])
        .map(
          (p) => `
            <li>${p.medicine} - ${p.dosage} - Qty: ${p.quantity}</li>
        `,
        )
        .join("");
      const patientLabel = getPatientDisplayName(visit.patient_id);
      const patient = getPatientDetails(visit.patient_id);
      const visitStatusLabel = visit.status
        ? visit.status.replaceAll("_", " ")
        : "-";

      return `
            <div class="visit-card">
                <div class="visit-header">
                    <div>
                        <h3 class="visit-title">Queue #${visit.queue_number}</h3>
                        <div class="visit-subtitle">${patientLabel}</div>
                <div class="visit-patient-grid">
                  <div><span>NIK</span><strong>${patient.nik || "-"}</strong></div>
                  <div><span>Gender</span><strong>${patient.gender || "-"}</strong></div>
                  <div><span>Umur</span><strong>${patient.age || "-"}</strong></div>
                  <div><span>Telepon</span><strong>${patient.phone_number || patient.phone || "-"}</strong></div>
                </div>
                        <div class="visit-meta">
                            <div><strong>Visit ID:</strong> ${visit.id}</div>
                            <div><strong>Symptoms:</strong> ${visit.symptoms}</div>
                            <div><strong>Diagnosis:</strong> ${visit.diagnosis || "-"}</div>
                        </div>
                    </div>
                    <span class="badge ${visit.status}">${visitStatusLabel}</span>
                </div>
                <div class="inline-row visit-actions">
                    ${visit.status === "registered" ? `<button type="button" class="btn btn-secondary small-btn" onclick="startExamineFromVisit('${visit.id}')">Periksa Visit Ini</button>` : ""}
                    ${visit.status === "waiting_pharmacy" ? `<button type="button" class="btn btn-secondary small-btn" onclick="startDispenseFromVisit('${visit.id}')">Proses Obat Visit Ini</button>` : ""}
                </div>
                ${prescriptions ? `<div><strong>Prescriptions:</strong><ul>${prescriptions}</ul></div>` : ""}
            </div>
        `;
    })
    .join("");
}

async function refreshDashboard() {
  try {
    const result = await apiFetch("/visits");
    const visits = result.data || [];

    animateCounter("total-visits", visits.length);
    animateCounter(
      "registered-visits",
      visits.filter((v) => v.status === "registered").length,
    );
    animateCounter(
      "waiting-visits",
      visits.filter((v) => v.status === "waiting_pharmacy").length,
    );
    animateCounter(
      "completed-visits",
      visits.filter((v) => v.status === "completed").length,
    );
  } catch {
    document.getElementById("total-visits").textContent = "0";
    document.getElementById("registered-visits").textContent = "0";
    document.getElementById("waiting-visits").textContent = "0";
    document.getElementById("completed-visits").textContent = "0";
  }
}

window.registerVisit = registerVisit;
window.addPrescriptionField = addPrescriptionField;
window.removePrescriptionField = removePrescriptionField;
window.examinePatient = examinePatient;
window.dispenseMedicine = dispenseMedicine;
window.refreshDashboard = refreshDashboard;
window.loadAllVisits = loadAllVisits;
window.loadVisitsByStatus = loadVisitsByStatus;
window.toggleCreatePatient = toggleCreatePatient;
window.searchPatient = searchPatient;
window.createPatient = createPatient;
window.selectPatient = selectPatient;
window.loadPatientsForPicker = loadPatientsForPicker;
window.loadVisitPickers = loadVisitPickers;
window.startExamineFromVisit = startExamineFromVisit;
window.startDispenseFromVisit = startDispenseFromVisit;
