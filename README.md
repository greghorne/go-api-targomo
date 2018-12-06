# go-api-targomo

Go: API that handles requests for isochrones from Targomo API.  The returned JSON is ready for use in LeafletJS.com

- This was written for using in a different project and the functionality is narrow in scope.
- The API returns JSON that make up the verticies of the isochrone (polygon).
- API requests are limited to North America.

__*Usage:*__ *http://myserver:8001/v1/targomo-isochrone/{lng}/{lat}/{time}/{key}*

- __*lng*__ => longitude (decimal degrees)
- __*lat*__ => latitude (decimal degrees)
- __*time*__ => drive time polygon in seconds
- __*key*__ => Targomo key

__*Example API Call & Return Value:*__

-   http://myserver:8001/v1/targomo-isochrone/-95.9668089440304/36.1329758/60/my_targomo_key
-   {"targomo":"[[36.13521042,-95.97184234],[36.13526263,-95.97179903],[36.13550974,-95.97149468],[36.13569273,-95.97114797],[36.13580458,-95.97077223],[36.13582497,-95.97055347],[36.13583302,-95.97055099],[36.13617797,-95.97036465],[36.13647994,-95.97011461],[36.13672733,-95.96981045],[36.13691063,-95.96946388],[36.1370228,-95.96908821],[36.13705952,-95.96869787],[36.13705694,-95.96810696],[36.13711197,-95.9680031],[36.13722439,-95.96762766],[36.13726145,-95.96723751],[36.13722171,-95.96684762],[36.13715258,-95.96662243],[36.13742975,-95.96654635],[36.13778032,-95.9663708],[36.1380899,-95.96613025],[36.1383466,-95.96583391],[36.13854057,-95.9654932],[36.13866433,-95.96512119],[36.13871315,-95.96473218],[36.13868513,-95.96434112],[36.13858137,-95.96396304],[36.13840584,-95.96361247],[36.1381653,-95.96330288],[36.13786897,-95.96304616],[36.13752826,-95.96285219],[36.13751237,-95.9628469],[36.13747247,-95.96271234],[36.13728985,-95.96236543],[36.13704306,-95.96206082],[36.13682553,-95.96187997],[36.13677876,-95.96179249],[36.13653001,-95.96148947],[36.13622692,-95.96124079],[36.13594055,-95.96108776],[36.13584642,-95.96092104],[36.13559667,-95.96062846],[36.13529552,-95.96038911],[36.13523705,-95.96035875],[36.13521374,-95.96013803],[36.1350973,-95.95976368],[36.13491006,-95.95941923],[36.13465923,-95.95911793],[36.13435444,-95.95887135],[36.1340074,-95.95868896],[36.13363145,-95.95857778],[36.13324104,-95.95854207],[36.13285116,-95.95858321],[36.13247679,-95.95869962],[36.13213233,-95.95888682],[36.131831,-95.95913763],[36.13158439,-95.9594424],[36.13140197,-95.95978942],[36.13129076,-95.96016536],[36.13126512,-95.96044538],[36.13118043,-95.9604935],[36.13088393,-95.96074992],[36.13064313,-95.96105925],[36.1304673,-95.9614096],[36.13036319,-95.96178753],[36.13033479,-95.9621785],[36.13038321,-95.9625675],[36.13050656,-95.96293958],[36.13070013,-95.96328046],[36.13095647,-95.96357703],[36.13126574,-95.96381791],[36.13127675,-95.96382344],[36.13127927,-95.96418949],[36.13128117,-95.96445591],[36.13128177,-95.96454656],[36.13128259,-95.96466426],[36.13128151,-95.96489839],[36.13126218,-95.96502989],[36.1312483,-95.96522758],[36.13122937,-95.96523195],[36.13087196,-95.96539304],[36.13055283,-95.96562075],[36.13028426,-95.96590635],[36.13007656,-95.96623885],[36.12993771,-95.96660548],[36.12987305,-95.96699215],[36.12988506,-95.967384],[36.12997329,-95.96776599],[36.13013433,-95.96812342],[36.13036201,-95.96844257],[36.13064757,-95.96871118],[36.13065497,-95.9687158],[36.13086759,-95.96901191],[36.13097298,-95.96911045],[36.13099685,-95.96914067],[36.13107725,-95.9692092],[36.13104463,-95.96926745],[36.13092349,-95.96964028],[36.1308774,-95.97002957],[36.13090813,-95.97042038],[36.1310145,-95.97079768],[36.13119243,-95.97114699],[36.13136221,-95.97136242],[36.13138195,-95.97154964],[36.13149833,-95.97192402],[36.1316855,-95.97226851],[36.13193628,-95.97256986],[36.13224103,-95.9728165],[36.13258804,-95.97299895],[36.13296397,-95.9731102],[36.13335439,-95.97314597],[36.13374428,-95.9731049],[36.13411867,-95.97298855],[36.13446317,-95.9728014],[36.13476454,-95.97255064],[36.1350112,-95.97224591],[36.13519368,-95.97189891],[36.13521042,-95.97184234]]"}

-   On error returns: {"targomo":"msg"}
-   msg = Targomo API call error message
