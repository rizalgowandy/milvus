import time
import pytest

from base.client_base import TestcaseBase
from common import common_func as cf
from common import common_type as ct
from common.common_type import CaseLabel

prefix = "e2e_"


class TestE2e(TestcaseBase):
    """ Test case of end to end"""
    @pytest.mark.tags(CaseLabel.L2)
    # @pytest.mark.parametrize("name", [(cf.gen_unique_str(prefix))])
    def test_milvus_default(self):
        from utils.util_log import test_log as log
        # create
        name = cf.gen_unique_str(prefix)
        t0 = time.time()
        collection_w = self.init_collection_wrap(name=name)
        tt = time.time() - t0
        log.debug(f"assert create: {tt}")
        assert collection_w.name == name

        # insert
        data = cf.gen_default_list_data()
        t0 = time.time()
        _, res = collection_w.insert(data)
        tt = time.time() - t0
        log.debug(f"assert insert: {tt}")
        assert res

        # flush
        t0 = time.time()
        assert collection_w.num_entities == len(data[0])
        tt = time.time() - t0
        log.debug(f"assert flush: {tt}")

        # search
        collection_w.load()
        search_vectors = cf.gen_vectors(1, ct.default_dim)
        search_params = {"metric_type": "L2", "params": {"nprobe": 16}}
        t0 = time.time()
        res_1, _ = collection_w.search(data=search_vectors,
                                       anns_field=ct.default_float_vec_field_name,
                                       param=search_params, limit=1)
        tt = time.time() - t0
        log.debug(f"assert search: {tt}")
        assert len(res_1) == 1
        collection_w.release()

        # index
        d = cf.gen_default_list_data(nb=2000)
        collection_w.insert(d)
        assert collection_w.num_entities == len(data[0]) + 2000
        _index_params = {"index_type": "IVF_SQ8", "params": {"nlist": 64}, "metric_type": "L2"}
        t0 = time.time()
        index, _ = collection_w.create_index(field_name=ct.default_float_vec_field_name,
                                             index_params=_index_params,
                                             name=cf.gen_unique_str())
        tt = time.time() - t0
        log.debug(f"assert index: {tt}")
        assert len(collection_w.indexes) == 1

        # search
        t0 = time.time()
        collection_w.load()
        tt = time.time() - t0
        log.debug(f"assert load: {tt}")
        search_vectors = cf.gen_vectors(1, ct.default_dim)
        t0 = time.time()
        res_1, _ = collection_w.search(data=search_vectors,
                                       anns_field=ct.default_float_vec_field_name,
                                       param=search_params, limit=1)
        tt = time.time() - t0
        log.debug(f"assert search: {tt}")

        # query
        term_expr = f'{ct.default_int64_field_name} in [1001,1201,4999,2999]'
        t0 = time.time()
        res, _ = collection_w.query(term_expr)
        tt = time.time() - t0
        log.debug(f"assert query result {len(res)}: {tt}")
        # assert len(res) == 4
