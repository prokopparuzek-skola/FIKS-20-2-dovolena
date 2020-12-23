# Dovolená

## Algoritmus
Algoritmus má dva kroky.\
1) Vytvoření párů jednostranně

V podstatě vezmu seznam mužů/žen a každému přiřadím nejbližší neobsazený protějšek. Má to jednu výjimku, pokud naleznu 
někoho, kdo má už obsazený protějšek, ale někým kdo ho má dále v seznamu než on, tak protějšky prohodím. Snažím se tak 
všem nalézt co nejoptimálnější protějšek.

Pokud naleznu, že někdo už mi protějšek sebral, ale má ho dříve, nebo stejně tak bych ho naštval, takže pokračuji. Pokud 
ho mám já dřív a seberu mu ho, tak ho sic naštvu, avšak zamezím situaci, kdybychom oba byli nespokojení a mohli si 
polepšit a to samé platí o protějšcích. Takhle by se mohlo stát, že já budu nespokojen on taky a protějšky též, protože 
on získal někoho kdo chce více mne. Takže tu je šance na více spokojených lidí za cenu 1 nespokojeného.\
2) Oboustranné vyrovnání

Porovnání obou seznamů chtěných protějšků. Projdu jeden seznam a porovnávám k němu druhý. Pokud naleznu shodu, chtějí se 
navzájem, je vše v pořádku. Pokud protějšek chce někoho jiného: už jsem ho spároval, tj. je spokojený, nebo jsem ho 
ještě nenapároval a pokud i on vidí vybranou ženu lépe tak je spáruji, jinak by nebyli spokojení a aktuálnímu jedinci 
dám ženu toho se kterým jsem ho prohodil, tím mi vznikne další spokojený pár.

## Složitost

V prvním kroku projdu všechny muže/ženy a u každého z nich projdu $N$ protějšků, dokud nenajdu ten pravý. To dává 
dohromady $O(N^2)$.\
V druhém kroku už procházím pouze $N$.\
V součtu to dá asymptoticky $O(N^2)$.
