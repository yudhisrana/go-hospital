const API_BASE = 'http://localhost:8080';

const state = {
    prescriptionCount: 0,
};

document.addEventListener('DOMContentLoaded', () => {
    initTabs();
    addPrescriptionField();
    refreshDashboard();
    loadAllVisits();
});

function initTabs() {
    const buttons = document.querySelectorAll('.tab-button');
    const tabs = document.querySelectorAll('.tab-content');

    buttons.forEach((button) => {
        button.addEventListener('click', () => {
            const target = button.dataset.tab;

            buttons.forEach((b) => b.classList.remove('active'));
            tabs.forEach((tab) => tab.classList.remove('active'));

            button.classList.add('active');
            document.getElementById(target).classList.add('active');
        });
    });
}

function formatResponse(data) {
    return JSON.stringify(data, null, 2);
}

async function apiFetch(path, options = {}) {
    const res = await fetch(`${API_BASE}${path}`, {
        headers: {
            'Content-Type': 'application/json',
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

function setBoxContent(selector, message, type = 'success') {
    const box = document.querySelector(selector);
    if (!box) return;
    box.className = `response-box ${type}`;
    box.textContent = typeof message === 'string' ? message : formatResponse(message);
}

function clearBox(selector) {
    const box = document.querySelector(selector);
    if (!box) return;
    box.className = 'response-box';
    box.textContent = '';
}

async function registerVisit() {
    const patientId = document.getElementById('patientId').value.trim();
    const symptoms = document.getElementById('symptoms').value.trim();

    if (!patientId || !symptoms) {
        setBoxContent('#registerResponse', 'ID pasien dan gejala wajib diisi.', 'error');
        return;
    }

    try {
        const result = await apiFetch('/visits/register', {
            method: 'POST',
            body: JSON.stringify({ patient_id: patientId, symptoms }),
        });
        setBoxContent('#registerResponse', result, 'success');
        document.getElementById('registerForm').reset();
        await refreshDashboard();
        await loadAllVisits();
    } catch (error) {
        setBoxContent('#registerResponse', error.message, 'error');
    }
}

// Toggle create patient box
function toggleCreatePatient() {
    const box = document.getElementById('createPatientBox');
    if (!box) return;
    box.style.display = box.style.display === 'none' ? 'block' : 'none';
}

// Search patient by NIK or name (client-side filtering)
async function searchPatient() {
    const q = document.getElementById('patientQuery').value.trim().toLowerCase();
    const resultBox = document.getElementById('patientSearchResult');
    resultBox.textContent = '';
    if (!q) {
        resultBox.textContent = 'Masukkan NIK atau nama untuk mencari.';
        return;
    }

    try {
        const res = await apiFetch('/patients');
        const patients = res.data || [];
        const matched = patients.filter(p => (p.nik && p.nik.toLowerCase().includes(q)) || (p.name && p.name.toLowerCase().includes(q)));
        if (matched.length === 0) {
            resultBox.textContent = 'Pasien tidak ditemukan.';
            return;
        }
        // if only one match, select it
        if (matched.length === 1) {
            document.getElementById('patientId').value = matched[0].id;
            resultBox.textContent = `Pasien dipilih: ${matched[0].name} (ID: ${matched[0].id})`;
            return;
        }

        // multiple matches: show clickable list
        resultBox.innerHTML = matched.map(p => `<div style="margin-bottom:6px;"><button class="btn btn-secondary small-btn" onclick="selectPatient('${p.id}','${p.name.replace(/'/g,"\'")}')">Pilih</button> &nbsp; ${p.name} — ${p.nik} — <small>${p.id}</small></div>`).join('');
    } catch (err) {
        resultBox.textContent = 'Gagal mencari pasien: ' + err.message;
    }
}

function selectPatient(id, name) {
    document.getElementById('patientId').value = id;
    document.getElementById('patientSearchResult').textContent = `Pasien dipilih: ${name} (ID: ${id})`;
}

// Create a new patient via API and auto-select
async function createPatient() {
    const nik = document.getElementById('new_nik').value.trim();
    const name = document.getElementById('new_name').value.trim();
    const age = Number(document.getElementById('new_age').value);
    const gender = document.getElementById('new_gender').value.trim();
    const birth_date = document.getElementById('new_birthdate').value.trim();
    const address = document.getElementById('new_address').value.trim();
    const phone_number = document.getElementById('new_phone').value.trim();

    if (!nik || !name || !age || !gender || !birth_date || !address || !phone_number) {
        setBoxContent('#registerResponse', 'Semua field pasien baru harus diisi.', 'error');
        return;
    }

    try {
        const res = await apiFetch('/patients', {
            method: 'POST',
            body: JSON.stringify({ nik, name, age, gender, birth_date, address, phone_number }),
        });
        const patient = res.data || res;
        // set patientId and close create box
        if (patient && patient.id) {
            document.getElementById('patientId').value = patient.id;
            setBoxContent('#registerResponse', `Pasien dibuat dan dipilih: ${patient.name} (ID: ${patient.id})`, 'success');
            document.getElementById('createPatientBox').style.display = 'none';
            document.getElementById('patientSearchResult').textContent = '';
        } else {
            setBoxContent('#registerResponse', 'Pasien dibuat, namun response tidak mengandung ID.', 'error');
        }
    } catch (err) {
        setBoxContent('#registerResponse', 'Gagal membuat pasien: ' + err.message, 'error');
    }
}

function addPrescriptionField() {
    state.prescriptionCount += 1;
    const wrapper = document.getElementById('prescriptionFields');
    const id = `prescription-${state.prescriptionCount}`;

    const row = document.createElement('div');
    row.className = 'prescription-item';
    row.id = id;
    row.innerHTML = `
        <input type="text" class="medicine-name" placeholder="Medicine" />
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
    const visitId = document.getElementById('visitId').value.trim();
    const diagnosis = document.getElementById('diagnosis').value.trim();
    const rows = document.querySelectorAll('#prescriptionFields .prescription-item');

    if (!visitId || !diagnosis) {
        setBoxContent('#examineResponse', 'Visit ID dan diagnosis wajib diisi.', 'error');
        return;
    }

    const prescriptions = [];
    rows.forEach((row) => {
        const medicine = row.querySelector('.medicine-name').value.trim();
        const dosage = row.querySelector('.medicine-dosage').value.trim();
        const quantity = Number(row.querySelector('.medicine-qty').value);

        if (medicine && dosage && quantity > 0) {
            prescriptions.push({ medicine, dosage, quantity });
        }
    });

    if (prescriptions.length === 0) {
        setBoxContent('#examineResponse', 'Minimal 1 resep harus diisi.', 'error');
        return;
    }

    try {
        const result = await apiFetch(`/visits/${visitId}/examine`, {
            method: 'PATCH',
            body: JSON.stringify({ diagnosis, prescriptions }),
        });
        setBoxContent('#examineResponse', result, 'success');
        document.getElementById('examineForm').reset();
        document.getElementById('prescriptionFields').innerHTML = '';
        state.prescriptionCount = 0;
        addPrescriptionField();
        await refreshDashboard();
        await loadAllVisits();
    } catch (error) {
        setBoxContent('#examineResponse', error.message, 'error');
    }
}

async function dispenseMedicine() {
    const visitId = document.getElementById('dispenseVisitId').value.trim();

    if (!visitId) {
        setBoxContent('#dispenseResponse', 'Visit ID wajib diisi.', 'error');
        return;
    }

    try {
        const result = await apiFetch(`/visits/${visitId}/dispense`, {
            method: 'PATCH',
        });
        setBoxContent('#dispenseResponse', result, 'success');
        document.getElementById('dispenseForm').reset();
        await refreshDashboard();
        await loadAllVisits();
    } catch (error) {
        setBoxContent('#dispenseResponse', error.message, 'error');
    }
}

async function loadAllVisits() {
    try {
        const result = await apiFetch('/visits');
        renderVisits(result.data || []);
    } catch (error) {
        const list = document.getElementById('visitsList');
        list.innerHTML = `<div class="response-box error">${error.message}</div>`;
    }
}

async function loadVisitsByStatus(status) {
    try {
        const result = await apiFetch(`/visits/status/${status}`);
        renderVisits(result.data || []);
    } catch (error) {
        const list = document.getElementById('visitsList');
        list.innerHTML = `<div class="response-box error">${error.message}</div>`;
    }
}

function renderVisits(visits) {
    const list = document.getElementById('visitsList');

    if (!visits || visits.length === 0) {
        list.innerHTML = '<div class="visit-card">Belum ada visit.</div>';
        return;
    }

    list.innerHTML = visits.map((visit) => {
        const prescriptions = (visit.prescriptions || []).map((p) => `
            <li>${p.medicine} - ${p.dosage} - Qty: ${p.quantity}</li>
        `).join('');

        return `
            <div class="visit-card">
                <div class="visit-header">
                    <div>
                        <h3 class="visit-title">Queue #${visit.queue_number} - ${visit.patient_id}</h3>
                        <div class="visit-meta">
                            <div><strong>Visit ID:</strong> ${visit.id}</div>
                            <div><strong>Symptoms:</strong> ${visit.symptoms}</div>
                            <div><strong>Diagnosis:</strong> ${visit.diagnosis || '-'}</div>
                        </div>
                    </div>
                    <span class="badge ${visit.status}">${visit.status}</span>
                </div>
                ${prescriptions ? `<div><strong>Prescriptions:</strong><ul>${prescriptions}</ul></div>` : ''}
            </div>
        `;
    }).join('');
}

async function refreshDashboard() {
    try {
        const result = await apiFetch('/visits');
        const visits = result.data || [];

        document.getElementById('total-visits').textContent = visits.length;
        document.getElementById('registered-visits').textContent = visits.filter(v => v.status === 'registered').length;
        document.getElementById('waiting-visits').textContent = visits.filter(v => v.status === 'waiting_pharmacy').length;
        document.getElementById('completed-visits').textContent = visits.filter(v => v.status === 'completed').length;
    } catch {
        document.getElementById('total-visits').textContent = '0';
        document.getElementById('registered-visits').textContent = '0';
        document.getElementById('waiting-visits').textContent = '0';
        document.getElementById('completed-visits').textContent = '0';
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

