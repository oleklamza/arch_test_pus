# Usługa analizy logów _logscope_

## Ogólny opis
Administratorom systemu sprzedażowego _TwójBazarek_ potrzebne jest narzędzie dostarczające raporty dotyczące wykorzystania ich usług, pojawiających się błędów i ogólnej kondycji systemu.

Usługa raportowania bazuje na pliku logów zapisywanym w formacie CSV z następującymi kolumnami:
- `timestamp` -- data i czas żądania (ISO 8601 w UTC), np. _2026-10-13T10:15:02_
- `ip` -- adres IP klienta, np. _192.168.1.13_
- `method` -- metoda HTTP, np. _GET_
- `endpoint` -- ścieżka żądania, np. _/api/products_
- `status` -- kod statusu HTTP, np. _200_
- `resp_time_ms` -- czas obsługi żądania w milisekundach, np. _76_

Przykładowa zawartość pliku*:
```
timestamp,           ip,            method, endpoint,      status, resp_time_ms
2026-10-13T08:00:09, 91.103.69.17,  GET,    /api/search ,  200,    162
2026-10-13T08:00:11, 185.68.207.12, GET,    /api/products, 500,    646
2026-10-13T08:00:26, 83.12.170.59,  POST,   /api/payments, 201,    386
```
_* Uwaga: spacje dodane dla zwiększenia czytelności; w plikach CSV ich nie ma._


## Kluczowe zagadnienia

**AO:**
- analiza wpływu decyzji architektonicznych na wydajność usługi
- optymalizacja przetwarzania danych i przechowywanie wyników w pamięci
- kompromis między wydajnością a aktualnością danych

**TO:**
- projektowanie i wykonywanie scenariuszy testów obciążeniowych w k6
- przegląd typów testów obciążeniowych
- analiza wyników testów: czas odpowiedzi, przepustowość, percentyl p95 i liczba błędów
- porównywanie wyników dla różnych poziomów i profili obciążenia
- wizualizacja wyników testów wydajnościowych

**PUS:**
- implementacja endpointu HTTP `/stats` w Go do analizy logów z pliku CSV
- parsowanie i przetwarzanie danych z wykorzystaniem standardowych bibliotek Go
- pomiar wydajności usługi i optymalizacja sposobu przetwarzania danych


## Pliki
- _log_1k.csv_ -- log z danymi (1000 wierszy)
- _log_10k.csv_ -- log z danymi (10 000 wierszy)
- _log_100k.csv_ -- log z danymi (100 000 wierszy)
