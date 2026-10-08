# System zapisu na zajęcia

Przykład spina tematykę z AO, TO i PUS.

## Ogólny opis
W ustalonym dniu, w samo południe, uruchamiane są zapisy na zajęcia w kolejnym semestrze. Każdy student może przeglądać ofertę, wybrać zajęcia i zapisać się do grupy.

Grupy mają ograniczoną liczbę miejsc.

W pierwszych minutach po uruchomieniu zapisów z systemu korzysta bardzo duża liczba studentów jednocześnie.

## Cel
**Jak zaprojektować i przetestować taki system, żeby w momencie największego obciążenia spełniał dobrze swoją funkcję?**

## Kluczowe zagadnienia

**AO:**
- analiza wymagań funkcjonalnych i niefunkcjonalnych
- atrybuty jakościowe
- problematyka obsługi współbieżności żądań

**TO:**
- scenariusze testów funkcjonalnych
- przypadki testowe testów niefunkcjonalnych
- wprowadzenie do testów obciążeniowych
- zapoznanie z narzędziem _k6_

**PUS:**
- budowa mocka usługi sieciowej w Go z wykorzystaniem standardowych bibliotek

## Pliki
- _enroll_mock_simple.go_ -- szablon podstawowego serwera HTTP
- _simulate_overload.go_ -- kod funkcji symulującej większe obciążenie serwera