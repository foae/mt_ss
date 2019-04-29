package main

import (
	"bytes"
	"io"
	"log"
	"math/rand"
	"unicode"
)

type tecst []byte

var blob = []byte(`
I

Un apartament bine aerisit, compus din trei încăperi principale, având terasă cu geamlâc și sonerie. În față, salonul somptuos, al cărui perete din fund este ocupat de o bibliotecă de stejar masiv, totdeauna strâns înfășurată în cearceafuri ude... O masă fără picioare, bazată pe calcule și probabilități, suportă un vas ce conține esența eternă a „lucrului în sine", un cățel de usturoi, o statuetă ce reprezintă un popă (ardelenesc) ținând în mână o sintaxă și... 20 de bani bacșiș... Restul nu prezintă nici o importanță. Trebuiește însă reținut că această cameră, vecinie pătrunsă de întuneric, nu are nici uși, nici ferestre și nu comunică cu exteriorul decât prin ajutorul unui tub, prin care uneori iese fum și prin care se poate vedea, în timpul nopții, cele șapte emisfere ale lui Ptolemeu, iar în timpul zilei doi oameni cum coboară din maimuță și un șir finit de bame uscate, alături de Auto-Kosmosul infinit și inutil...

A doua încăpere, care formează un interior turc, este decorată cu mult fast și conține tot ceea ce luxul oriental are mai rar și mai fantastic... Nenumărate covoare de preț, sute de arme vechi, încă pătate de sânge eroic, căptușesc colonadele sălii, iar imenșii ei pereți sunt, conform obiceiului oriental, sulemeniți în fiecare dimineață, alteori măsurați, între timp, cu compasul pentru a nu scădea la întâmplare.

De aci, printr-o trapă făcută anume în dușumea, se ajunge, din partea stângă, în o subt-pământă ce formează sala de recepție, iar din partea dreaptă, prin ajutorul unui cărucior pus în mișcare cu manivela, se pătrunde într-un canal răcoros, al căruia unul din capete nu se știe unde se termină, iar celălalt, la partea opusă, într-o încăpere scundă, cu pământ pe jos și în mijlocul căreia se află bătut un țăruș, de care se află legată întreaga familie Stamate...


II

Acest om demn, unsuros și de formă aproape eliptică, din cauza nervozității excesive la care a ajuns de pe urma ocupațiilor ce le avea în consiliul comunal, este silit să mestece, mai toată ziua, celuloid brut, pe care apoi îl dă afară, fărâmițit și insalivat, asupra unicului său copil, gras, blazat și în etate de patru ani, numit Bufty... Micul băiat, din prea multă pietate filială, prefăcându-se însă că nu observă nimic, târăște o mică targa, pe uscat, în vreme ce mama sa, soția tunsă și legitimă a lui Stamate, ia parte la bucuria comună, compunând madrigale, semnate prin punere de deget.

Aceste ocupațiuni îndeajuns de obositoare îi fac, cu drept cuvânt, să se amuzeze, și atunci, ajungând uneori cu îndrăzneala până la inconștiență, se uită tustrei cu benoclul, printr-o spărtură a canalului, în Nirvana (care se află situată în aceeași circumscripție cu dânșii, începând lângă băcănia din colț), și aruncă în ea cu cocoloașe făcute din miez de pâine sau cu coceni de porumb. Alteori, pătrund în sala de recepție și dau drumul unor robinete expres construite acolo, până ce apa, revărsându-se, le-a ajuns în dreptul ochilor, când cu toții trag atunci, de bucurie, focuri de pistol în aer.

În ce privește personal pe Stamate, o ocupație care îl preocupă în gradul cel mai înalt este ca să ia seara, prin biserici, instantanee de pe sfinții mai în vârstă, pe cari le vinde apoi cu preț redus credulei sale soții și mai ales copilului Bufty, care are avere personală. Acest negoț nepermis nu l-ar fi exercitat pentru nimic în lume Stamate dacă nu ar fi dus lipsă aproape completă de mijloace, fiind silit chiar să facă armata când era abia în vârstă de un an, numai ca să poată ajuta, cât de curând, pe doi frățiori nevoiași ai săi, cu șoldurile scoase prea mult în afară, cauză pentru care fuseseră dați afară din slujbă.

Într-una din zile, lui Stamate, ocupat fiind cu obișnuitele sale cercetări filozofice, i se păru, o clipită, că a pus mâna și pe cealaltă jumătate a „lucrului în sine", când fu distras de o voce femeiască, o voce de sirenă, ce mergea drept la inimă și se auzea în depărtare, pierzându-se ca un ecou.

Alergând de urgență la tubul de comunicație, Stamate, spre marea lui înmărmurire, văzu cum, în aerul cald și îmbălsămat al serii, o sirenă cu gesturi și voce seducătoare își întindea corpul lasciv pe nisipul fierbinte al mării... în luptă puternică cu sine, pentru a putea să nu cadă pradă tentației, Stamate închirie atunci în grabă o corabie și, pornind în larg, își astupă urechile cu ceară împreună cu toți matrozii...

III

Sirena deveni însă tot mai provocatoare... Ea îl urmări pe întinsul apelor, cu cântări și gesturi perverse, până ce o duzină de Driade, Nereide și Tritoni avură tot timpul să se adune din larguri și adâncuri și să aducă pe o superbă cochilie de sidef o inocentă și decentă pâlnie ruginită.

Planul de seducțiune al seriosului și castului filozof putea fi astfel considerat ca reușit. Abia avu el timpul să se furișeze la tubul de comunicație, când zânele mării îi și depuseră, grațios, pâlnia în preajma locuinței sale, apoi, ușoare, zglobii, în râsete și chiote nebune, dispărură cu toate pe întinsul apelor.

Confuz, înnebunit, dezagregat, Stamate abia putu să apară cu căruciorul prin canal... Fără a pierde însă cu totul sângele rece, azvârli el de câteva ori cu țărâna asupra pâlniei și, după ce se ospăta cu puțină fiertură de ștevie, se aruncă, din diplomație, cu fața la pământ, rămânând astfel în nesimțire timp de opt zile libere, termenul necesar ce, după procedura civilă, credea dânsul că trebuia să treacă pentru a putea fi pus în posesiunea obiectului.

După această trecere de timp, reluându-și ocupațiunile cotidiane și pozițiunea verticală, Stamate se simți cu totul renăscut. Niciodată nu cunoscuse el până atunci divinii fiori ai dragostei. Se simțea acum mai bun, mai îngăduitor, și turburarea ce o încerca la vederea acestei pâlnii îl făcea să se bucure și totodată să sufere și să plângă ca un copil...

O scutură cu un otrep și, după ce îi unse găurile mai principale cu tinctură de iod, o luă cu sine și, cu legături de flori și dantele, o fixă alături și paralel cu tubul de comunicație, și, tot atunci, pentru prima oară, istovit de emoție, trecu printr-însa, ca fulgerul, și îi fură o sărutare.

Pentru Stamate, pâlnia deveni de atunci un simbol. Era singura ființă de sex femeiesc cu un tub de comunicație ce i-ar fi permis să satisfacă și cerințele dragostei, și interesele superioare ale științei. Uitându-și cu totul sacrele îndatoriri de tată și de soț, Stamate începu să-și taie în fiecare noapte, cu foarfecă, legăturile ce-l țineau atașat de țăruș și, spre a putea da frâu liber dragostei sale nețărmurite, începu să treacă din ce în ce mai des prin interiorul pâlniei, făcându-și vânt în ea de pe o trambulină construită expres și coborându-se apoi în mâini, cu o iuțeală vertiginoasă, pe o scară mobilă de lemn,la capătul căreia își rezuma rezultatul observărilor sale în afară.


IV

Fericirile mari sunt totdeauna de scurtă durată... într-una din nopți, Stamate, venind spre a-și face obișnuita-i datorie sentimentală, constată cu uimire și dezamăgire că, din cauze încă nepătrunse, orificiul de ieșire al pâlniei se strâmtase într-atâta, încât orice comunicație prin el era imposibilă. Nedumerit și totuși bănuitor, se puse la pândă, și a doua noapte, necrezându-și ochilor, văzu cu groază cum Bufty, urcat sus, gâfâind, fusese lăsat să intre și să treacă. Pietrificat, Stamate abia avu puterea să se ducă să se lege singur la țăruș; a doua zi însă luă o hotărâre supremă.

Mai întâi își îmbrățișa soția devotată și, după ce-i dădu în grabă o vopsea, o cusu într-un sac impermeabil, în scopul de a păstra mai departe, intactă, tradițiunea culturală a familiei. După aceea, în mijlocul unei nopți reci și întunecoase, luă el pâlnia și pe Bufty și, aruncându-i într-un tramvai, ce tocmai trecea la întâmplare, le făcu vânt cu dispreț în Nirvana. Totuși, mai târziu, sentimentul patern învinse, și Stamate, grație calculeleior și combinațiilor sale chimice, reuși să facă cu timpul, prin puterea științei, ca Bufty să ocupe acolo un post de subșef de birou.

Cât despre eroul nostru, Stamate, pentru ultima oară cătând prin tubul de comunicație, mai privi o dată Kosmosul cu ironie și indulgență. Suindu-se apoi pentru totdeauna în căruciorul cu manivelă, luă direcția spre capătul misterios al canalului și, mișcând manivela cu o stăruință crescândă, aleargă și astăzi, nebun, micșorându-și mereu volumul, cu scopul de a putea odată pătrunde și dispărea în infinitul mic.

Algazy nu vorbește nici o limbă europeană... Dacă însă îl aștepți în zori de zi, în faptul dimineței, și îi zici: „Bună ziua Algazy!" insistând mai mult pe sunetul z, Algazy zâmbește, iar spre a-și manifesta gratitudinea, bagă mâna în buzunar și trage de capătul unei sfori, făcând să-i tresalte de bucurie barba un sfert de oră... Deșurupat, grătarul îi servea să rezolve orice probleme mai grele, referitoare la curățirea și liniștea casei...

Algazy nu ia mită... O singură dată s-a pretat la o asemenea faptă, când era copist la Casa bisericii; dar nu a luat atunci bani, ci numai câteva cioburi de străchini, din dorința de a face dotă unor surori ale sale sărace, cari trebuiau să se mărite toate a doua zi...

Cea mai mare plăcere a lui Algazy - în afară de obișnuitele-i ocupații la prăvălie - este să se înhame de bunăvoie la o roabă și urmat cam la doi metri de coasociatul său Grummer - să alerge, în goana mare, prin praf și arșița soarelui, străbătând comunele rurale, în scopul unic de a aduna cârpe vechi, tinichele de untdelemn găurite și în special arșice, pe cari apoi le mănâncă împreună, după miezul nopții, în tăcerea cea mai sinistră...

Grummer are și un cioc de lemn aromatic...

Fire închisă și temperament bilios, stă toată ziua lungit sub tejghea, cu ciocul vârât într-o gaură sub podea...

Cum intri la ei în magazin, un miros delicios îți gâdilă nările... Ești întâmpinat la scară de un băiat cinstit, care, pe cap, în loc de păr, are fire de arnici verde; apoi ești salutat cu multă amabilitate de Algazy și poftit să stai jos pe un taburel.

Grummer stă și pândește... Perfid, cu privirea piezișă, scoțând mai întâi numai ciocul, pe care ostentativ îl prelinge în sus și în jos pe un jgheab anume făcut la muchea tejghelei, apare la urmă în întregime... Face prin tot felul de manopere pe Algazy sâi părăsească localul, apoi, insinuant, te atrage pe nesimțite în tot soiul de discuții - mai ales de sport și literatură - până ce, când îi vine bine, te plesnește de două ori cu ciocul peste burtă, de te face să alergi afară în stradă, urlând de durere.

Algazy, care are mai totdeauna neplăceri și discuții cu clienții, din cauza acestui procedeu nepermis al lui Grummer, iese în goană după tine, te poftește înapoi și, spre a-ți lua meritata satisfacție, îți dă dreptul - dacă ai cumpărat un obiect mai scump ca 15 bani - să... miroși puțin ciocul lui Grummer și, dacă vrei, să-l strângi cât de tare de o bășică cenușie de cauciuc, pe care o are înșurupată la spate, puțin deasupra feselor, ceea ce îl face să sară prin magazin fără să miște din genunchi, scoțând și sunete nearticulate...

Într-una din zile, Grummer, fără a-l anunța pe Algazy, luă roaba și porni singur în căutare de cârpe și arșice, dar la înapoiere, găsind din întâmplare și câteva resturi de poeme se prefăcu bolnav și, sub plapomă, le mâncă singur pe furiș. Algazy, simțind, intră după el acolo cu intenția sinceră de a-i face numai o ușoară morală, dar cu groază observă în stomacul lui Grummer că tot ce rămăsese bun din literatură fusese consumat și digerat.

Lipsit astfel pe viitor de orice hrană a lui mai aleasă, Algazy, drept compensație, mâncă toată bășica lui Grummer, în timp ce acesta dormea...

Dezesperat, a doua zi, Grummer - rămas fără bășică singur pe lume - luă pe bătrân în cioc și, după apusul soarelui, îl urcă cu furie pe vârful unui munte înalt... O luptă uriașă se încinse acolo între ei și ținu toată noaptea, până când, înspre ziuă, Grummer, învins, se oferi să restituie toată literatura înghițită. El o vomită în mâinile lui Algazy... Dar bătrânul, în pântecul căruia fermenții bășicii înghițite începuse să trezească fiorii literaturii viitorului, găsi că tot ce i se oferă este prea puțin și învechit...

Înfometați și nefiind în stare să găsească prin întunerec hrana ideală de care amândoi aveau atâta nevoie, reluară atunci lupta cu puteri îndoite și, sub pretext că se gustă numai pentru a se complecta și cunoaște mai bine, începură să se muște cu furie mereu crescândă, până ce, consumându-se treptat unul pe altul, ajunseră ambii la ultimul os... Algazy termină mai întâi...

Și iarăși, cine dintre noi se mai poate plânge că forța primordială, cauza cauzelor, nu poate fi niciodată atinsă, descoperită, când toți se căznesc să o apuce de la început, înapoi, și nimeni nu s-a încercat să o învăluie, pentru clipa de față, să o prindă măcar o dată pe flanc?

Și care e rostul să ții morțiș să descoperi vreo cauză, și că numai una singură și cea dintâi, când toate cauzele, din nenorocire, sunt și efecte și dau din ele efecte îndrăcit de multiple și de încâlcite.

Deci la ce bun să vrei numai o singură cauză, o forță inițială care vrem (trebuie) să fie și generatoare, când ea însăși ține cu încăpățânare să dea din ea numai multiplicitate; are setea ultimilor, a încâlcelei și contradicției; îi trebuie multe milioane de oameni, de muște, de bureți, de jivine, de astre, și aceasta încă cu prețul suferințelor lor. Îi trebuie și „peștele-cufăr", și peștele-fierăstrău" și are setea numărului, a distanțelor și idelilor mari, fără rost și necesitate...

La început - ziseră toți comesenii laolaltă - nu este adevărat că: „Cuvântul a fost la Dumnezeu și că D-zeu a fost cuvântul". La început - afirmară ei cu tărie - mai înainte de a fi fost ori „cuvânt" a fost „alfabetul surdo-mut", căci nu este probabil ca materia cosmică, astrele să fi învățat chiar de la început a g(?) ceva; că e prea posibil ca ele să nu fi fost în stare, la început, nici măcar să se ceară afară, și nici chiar să zică „papa" sau „mamă". Asemenea, este foarte probabil - ziseră mai departe comesenii - că corpii cerești s-au format nici din dărnicia lui D-zeu, nici din pofta lor proprie de a se învârti și a se face numai pentru atâta lucru din nimic ceva, apoi din gazoase solid ci e prea posibil ca ele să nu fi fost nici create, nici necreate, copii ai nimănui ieșiți din calcule reușite sau greșite și, cu chiu cu vai, în rate și hrăniți insuficient la institutul „Maternitatea cerească" cu lapte contrafăcut cu apă gazoasă de lăptăresele Căii-lactee.

Că admițând că ele se învârtesc numai din propriul lor gust, apoi e greu de presupus că fac aceasta în mod cu totul dezinteresat, fără nici o intenție cât de mică de procopseală pe setea mulțimilor și a distanțelor mari, fără rost și necesitate. Parcă ar fi chiar puțin comic să te învârtești în veci de veci gratuit și numai pentru a te vedea alții...

- Cum, interese meschine, egoiste la corpii cerești? protestă plebea ideologică, naivă, care aștepta rezultatul afară în curte.

Și totuși ei aveau și nu dreptate să se teamă astfel...

În adevăr, cine a putut mai întâi obliga materia și forța cosmică să fie ceva, când ele însele, la rândul lor, desființându-se, dându-și demisia, ar putea oricând obliga pe acel „cine" să nu fie „nimic"?!...
`)

func generate(s []byte) []byte {
	found := bytes.NewBuffer(make([]byte, 0))
	dflt := []byte(`Sincer, nu am nimic mai bun de zis în această privință, mi-a depășit literalmente așteptările.`)
	max := len(s) - 1

	w := s[rand.Intn(max):max]
	buf := bytes.NewBuffer(w)
	var firstDot, secondDot bool
	for {
		ch, _, err := buf.ReadRune()
		switch {
		case err == io.EOF:
			break
		case err != nil:
			log.Printf("reading yielded and error: %v", err)
			break
		}

		if firstDot && secondDot {
			break
		}

		// Avoid space after dot . New sentence
		if firstDot && unicode.IsSpace(ch) {
			continue
		} else {
			// Capture desired runes.
			found.WriteRune(ch)
		}

		// Continue until we have found
		// our first punctuation mark.
		if firstDot == false && unicode.IsSpace(ch) {
			firstDot = true
			continue
		}

	}

	if found.Len() == 0 {
		return dflt
	}

	return found.Bytes()
}

func (t *tecst) String() string {
	if t != nil {
		return string(*t)
	}

	return ""
}
