# Nazwa systemu
System rezerwacji przestrzeni biurowych _InvadeSpaces_

# Cel
Umożliwienie rezerwowania sal/przestrzeni biurowych na określone terminy oraz zarządzania utworzonymi rezerwacjami.

# Zakres funkcjonalny
- tworzenie rezerwacji
- sprawdzanie dostępności sal
- potwierdzanie rezerwacji
- anulowanie rezerwacji
- przeglądanie informacji o rezerwacji
- ...

**Poza zakresem:**
- płatności
- integracja z systemem kontroli dostępu do budynku
- automatyczne przydzielanie sal na podstawie planu zajęć
- wysyłanie powiadomień e-mail
- ...

# Aktorzy i słownik pojęć
___
W tej sekcji należy zebrać omówienie aktorów (_aktor -- odpowiedzialność_)
oraz definicje stosowanych pojęć (_pojęcie -- definicja_).  
___

## Aktorzy
| Aktor                         | Odpowiedzialność                   |
|-------------------------------|------------------------------------|
| gość                      | przeglądanie kalendarza rezerwacji |
| użytkownik zarejestrowany | ... |
| ...                           | ... |

## Słownik
| Pojęcie    | Definicja                                                    |
|------------|--------------------------------------------------------------|
| sala       | zasób, który można zarezerwować                              |
| rezerwacja | powiązanie użytkownika, terminu oraz sali |
| ...        | ... |

## Uszczegółowienie pojęć

### Stany rezerwacji

Rezerwacja może się znaleźć w jednym z czterech stanów:
- `PENDING` -- oczekuje na potwierdzenie
- `CONFIRMED` -- potwierdzona
- `CANCELED` -- anulowana
- `COMPLETED` -- zrealizowana


# Reguły biznesowe
___
Jest to kluczowy element specyfikacji będący punktem odniesienia dla implementacji i testów.
 
Schemat budowy reguły biznesowej:
- **identyfikator** -- pozwala jednoznacznie wskazać regułę, np. `BR-01`
- **nazwa** -- krótkie określenie jej znaczenia
- **opis** -- jednoznacznie sformułowane wymagania; opis można podzielić na
_warunki obowiązywania_ (czyli kiedy reguła ma zastosowanie) oraz _oczekiwane zachowanie_
(czyli co system powinien zrobić).

**Przykład:**

| BR-01       | Termin rezerwacji                                                             |
|-------------|-------------------------------------------------------------------------------|
| opis        | rezerwacja może się rozpoczynać wyłącznie w przyszłości                       |
| warunek     | użytkownik tworzy nową rezerwację                                             |
| zachowanie  | jeżeli termin jest wcześniejszy niż bieżący czas, rezerwacja zostaje odrzucona |
___
...


# Operacje systemu
___
 Ta sekcja ma zawierać opis operacji wykonywanych przez system.

Schemat budowy opisu operacji:
- **identyfikator** -- pozwala jednoznacznie wskazać operację, np. `OP-01`
- **cel** -- krótkie określenie celu operacji
- **wejście** -- dane, które należy dostarczyć
- **warunki wstępne** -- warunki, których spełnienie jest konieczne do przeprowadzenia operacji
- **reguły biznesowe** -- reguły związane z daną operacją
- **wynik pozytywny** -- stan systemu po wykonaniu operacji
- **możliwe błędy** -- potencjalne przyczyny odrzucenia wykonania operacji

**Przykład:**

| OP-01           | Utworzenie rezerwacji                         |
|-----------------|-----------------------------------------------|
| cel             | utworzenie nowej rezerwacji                   |
| wejście         | identyfikator sali, czas rozpoczęcia, czas trwania |
| warunki wstępne | użytkownik ma uprawnienia do rezerwowania sal |
| reguły biznesowe | spełnienie reguł BR-01, ...                   |
| wynik pozytywny | rezerwacja zostaje utworzona w stanie PENDING |
| możliwe błędy   | niepoprawne dane, brak uprawnień, konflikt terminów |

Minimalny zestaw operacji:
- `OP-01`: utworzenie rezerwacji
- `OP-02`: pobranie rezerwacji
- `OP-03`: potwierdzenie rezerwacji
- `OP-04`: anulowanie rezerwacji
- `OP-05`: sprawdzenie dostępności
___
...

# Kryteria akceptacji
___
Ta sekcja jest ściśle powiązana z testowaniem. Kryteria akceptacji pozwalają
na zweryfikowanie, czy system realizuje to, co zostało opisane w wymaganiach.

Proponowany format to **GWT** (_Given-When-Then_):
- _given_ -- stan początkowy, warunki
- _when_ -- wykonywane operacje
- _then_ -- oczekiwany rezultat

**Przykład:**

| AC-01 | Poprawne utworzenie rezerwacji |
|-------| ------------------------------ |
| G     | użytkownik jest uprawniony, sala jest dostępna, a dane rezerwacji są poprawne |
| W     | użytkownik wysyła żądanie utworzenia rezerwacji |
| T     | system tworzy rezerwację w stanie PENDING i zwraca jej identyfikator |
___