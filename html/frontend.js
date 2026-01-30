// Shared frontend utilities for Simple ERP test UI

async function apiFetch(path, options = {}) {
  const baseOptions = {
    headers: {
      "Accept": "application/json",
      ...(options.body ? { "Content-Type": "application/json" } : {})
    },
    ...options
  };

  const res = await fetch(path, baseOptions);
  let data = null;
  try {
    data = await res.json();
  } catch (_) {
    // ignore JSON parse errors, keep data as null
  }

  if (!res.ok) {
    const message =
      (data && (data.error || data.message)) ||
      `Request failed with status ${res.status}`;
    const err = new Error(message);
    err.status = res.status;
    err.data = data;
    throw err;
  }

  return data;
}

function clearElement(el) {
  while (el.firstChild) el.removeChild(el.firstChild);
}

function showMessage(container, type, text) {
  if (!container) return;
  clearElement(container);
  const div = document.createElement("div");
  div.textContent = text;
  div.style.padding = "0.5rem 0.75rem";
  div.style.borderRadius = "4px";
  div.style.fontSize = "0.9rem";
  if (type === "error") {
    div.style.background = "#fee2e2";
    div.style.color = "#b91c1c";
    div.style.border = "1px solid #fecaca";
  } else {
    div.style.background = "#dcfce7";
    div.style.color = "#166534";
    div.style.border = "1px solid #bbf7d0";
  }
  container.appendChild(div);
}

function renderTable(container, rows) {
  if (!container) return;
  clearElement(container);

  if (!rows || rows.length === 0) {
    const p = document.createElement("p");
    p.textContent = "No data.";
    p.style.color = "#6b7280";
    p.style.fontSize = "0.9rem";
    container.appendChild(p);
    return;
  }

  const table = document.createElement("table");
  table.style.borderCollapse = "collapse";
  table.style.width = "100%";
  table.style.marginTop = "0.75rem";
  table.style.fontSize = "0.85rem";

  const thead = document.createElement("thead");
  const headerRow = document.createElement("tr");

  const keys = Object.keys(rows[0]);
  keys.forEach((key) => {
    const th = document.createElement("th");
    th.textContent = key;
    th.style.textAlign = "left";
    th.style.padding = "0.4rem 0.5rem";
    th.style.borderBottom = "1px solid #e5e7eb";
    th.style.background = "#f9fafb";
    th.style.position = "sticky";
    th.style.top = "0";
    headerRow.appendChild(th);
  });
  thead.appendChild(headerRow);

  const tbody = document.createElement("tbody");
  rows.forEach((row) => {
    const tr = document.createElement("tr");
    keys.forEach((key) => {
      const td = document.createElement("td");
      const value = row[key];
      td.textContent = value === null || value === undefined ? "" : String(value);
      td.style.padding = "0.35rem 0.5rem";
      td.style.borderBottom = "1px solid #f3f4f6";
      tr.appendChild(td);
    });
    tbody.appendChild(tr);
  });

  table.appendChild(thead);
  table.appendChild(tbody);
  container.appendChild(table);
}
